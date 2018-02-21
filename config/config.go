/*
 * irrenhaus-gui, gtk client for irrenhaus.dyndns.dk
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

package config

import (
	"encoding/json"
	"os"

	api "github.com/fuchsi/irrenhaus-api"
)

// Configuration structure
type Configuration struct {
	Username string
	Password string
	Pin      string
	URL      string
}

// LoadConfig loads the configuration file
// It returns the parsed configration struct and any error encountered
func LoadConfig(configFile string) (Configuration, error) {
	file, err := os.Open(configFile)
	defer file.Close()
	if err != nil {
		return Configuration{}, err
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}

// DumpConfig writes the configuration file
// It returns any error encountered.
func DumpConfig(config Configuration, configFile string) error {
	file, err := os.Create(configFile)
	defer file.Close()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return err
	}

	return nil
}

// DumpCookies dumps the cookies for later reuse
// It returns any error encountered.
func DumpCookies(configPath string, cookies api.Cookies) error {
	file, err := os.Create(configPath + "cookies.json")
	defer file.Close()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cookies)
	if err != nil {
		return err
	}

	return nil
}

// LoadCookies loads the cookies
// It returns the parsed cookie struct and any error encountered
func LoadCookies(configPath string) (api.Cookies, error) {
	file, err := os.Open(configPath + "cookies.json")
	defer file.Close()
	if err != nil {
		return api.Cookies{}, err
	}
	decoder := json.NewDecoder(file)
	cookies := api.Cookies{}
	err = decoder.Decode(&cookies)
	if err != nil {
		return cookies, err
	}

	return cookies, nil
}
