# vocx

[![CircleCI](https://circleci.com/gh/martinrue/vocx.svg?style=svg)](https://circleci.com/gh/martinrue/vocx)

`vocx` transcribes Esperanto text into phonetic Polish for use in professional TTS engines.

## Background

Commercial TTS engines tend not to support minority languages, particularly constructed languages such as Esperanto. It turns out Esperanto shares lots of sounds with Polish. By transcribing Esperanto to Polish, we can make commercial TTS engines give us a good approximation for spoken Esperanto.

## Usage

```go
t := vocx.NewTranscriber()
t.Transcribe("Ĉu vi ŝatas Esperanton? Esperanto estas facila lingvo.")

// czu wij szatas esperanton? esperanto estas fatssila lijngwo.
```

### Custom Rules

To override the default rules used during transcriptions, call the `LoadRules` function, passing it a custom rule document. See [default_rules.go](./default_rules.go) for the correct structure.
