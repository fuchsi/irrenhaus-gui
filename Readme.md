# irrenhaus-gui

A GTK client for the irrenhaus tracker

## Precompiled versions

- Linux: [irrenhaus-gui-linux-amd64](https://github.com/fuchsi/irrenhaus-gui/releases/download/v0.1.0/irrenhaus-gui-linux-amd64)
- Windows: *coming soon™*

## Dependencies / System Requirements

You need a recent version of GTK 2.x

## Compiling

### Linux

You will need Go and the development files for GTK 2.x. Both should be available from your package manager.

Create a directory that contains the directories `bin`, `src` and `pkg` and set the `GOPATH` environment variable to that directory.

Then compiling is done by running

```bash
git clone https://github.com/fuchsi/irrenhaus-gui
cd irrenhaus-gui
make linux
make install
```

or

```bash
go get github.com/fuchsi/irrenhaus-gui
```
