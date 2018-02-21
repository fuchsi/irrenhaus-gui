/*
 * irrenhaus-gui, GTK client for irrenhaus.dyndns.dk
 * Copyright (C) 2018  Daniel Müller
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>
 */

package maketorrent

import (
	"bufio"
	"crypto/sha1"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"strings"

	"github.com/fuchsi/torrentfile"
)

type CreateProgress struct {
	NumPieces    uint64
	HashedPieces uint64
	Finished     bool
}

func CreateFromSingleFile(filename string, pieceLength uint64, cProgress chan CreateProgress) (torrentfile.File, [][torrentfile.PIECE_SIZE]byte) {
	finfo, _ := os.Stat(filename)
	file := torrentfile.File{Path: finfo.Name(), Length: uint64(finfo.Size())}
	numPieces := numPieces(file.Length, pieceLength)

	pieces := make([][torrentfile.PIECE_SIZE]byte, numPieces)
	progress := CreateProgress{NumPieces: numPieces, HashedPieces: 0}

	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	pieceIndex := 0

	for {
		buf := make([]byte, pieceLength)
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			break
		}
		if n < int(pieceLength) {
			pieceBuf := make([]byte, n)
			copy(pieceBuf, buf)
			buf = pieceBuf
		}
		pieces[pieceIndex] = sha1.Sum(buf)
		pieceIndex++
		progress.HashedPieces++
		cProgress <- progress
	}

	progress.Finished = true
	cProgress <- progress

	return file, pieces
}

func CreateFromDirectory(filename string, pieceLength uint64, cProgress chan CreateProgress) ([]torrentfile.File, [][torrentfile.PIECE_SIZE]byte) {
	filename = strings.TrimRight(filename, "/")
	files := collectFiles(filename)
	var totalSize uint64

	for i, f := range files {
		files[i].Path = strings.TrimPrefix(f.Path, filename+"/") // there must be a better way to alter the path
		totalSize += f.Length
	}

	numPieces := numPieces(totalSize, pieceLength)

	pieces := make([][torrentfile.PIECE_SIZE]byte, numPieces)

	bufLen := pieceLength
	pieceIndex := 0
	pieceBuf := make([]byte, pieceLength)
	off := uint64(0)
	c := make(chan piece, runtime.GOMAXPROCS(0))
	progress := CreateProgress{NumPieces: numPieces, HashedPieces: 0}

	for _, f := range files {
		fp, err := os.Open(filename + "/" + f.Path)
		if err != nil {
			fp.Close()
			log.Fatal(err)
		}
		reader := bufio.NewReader(fp)

		for {
			buf := make([]byte, bufLen)
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			} else if err == io.EOF {
				break
			}
			length := uint64(n)

			if length < bufLen { // got less bytes than pieceLen from file (reached EOF while reading)
				bufLen = pieceLength - length
				copy(pieceBuf[off:], buf[:length]) // copy length bytes from buf to pieceBuf
				off = length                       // set new offset for pieceBuf to length
			} else if off != 0 { // got the remaining bytes from the next file
				copy(pieceBuf[off:], buf) // copy remaining bytes from buf to pieceBuf
				off = 0                   // reset offset and bufLen
				bufLen = pieceLength
			} else { // normal operation, just copy buf to pieceBuf
				copy(pieceBuf, buf)
			}

			if off == 0 { // hash the piece if offset is zero
				go buildHash(piece{index: pieceIndex}, pieceBuf, c)
				progress.HashedPieces++
				cProgress <- progress
				//pieces[pieceIndex] = sha1.Sum(pieceBuf)
				pieceIndex++
			}
		}

		fp.Close()
	}

	// add remaining bytes from buffer
	if off != 0 {
		//pieces[pieceIndex] = sha1.Sum(pieceBuf[:off])
		go buildHash(piece{index: pieceIndex}, pieceBuf, c)
		progress.HashedPieces++
		cProgress <- progress
	}

	for i := uint64(0); i < numPieces; i++ {
		p := <-c
		pieces[p.index] = p.hash

	}

	progress.Finished = true
	cProgress <- progress

	return files, pieces
}

type piece struct {
	index int
	hash  [torrentfile.PIECE_SIZE]byte
}

func buildHash(p piece, data []byte, c chan piece) {
	p.hash = sha1.Sum(data)
	c <- p
}

func collectFiles(filename string) []torrentfile.File {
	var filelist []torrentfile.File
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			for _, inner := range collectFiles(filename + "/" + f.Name()) {
				filelist = append(filelist, inner)
			}
		} else {
			filelist = append(filelist, torrentfile.File{Length: uint64(f.Size()), Path: filename + "/" + f.Name()})
		}
	}

	return filelist
}

func numPieces(filesize, pieceLength uint64) uint64 {
	return uint64(math.Ceil(float64(filesize) / float64(pieceLength)))
}
