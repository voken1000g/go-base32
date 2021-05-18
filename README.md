# Voken flavored Base32 in GoLang

Compute Voken flavored Base32 encoding/decoding, with auto checksum.

## API

| Function         | Params   | Returns           |
| ---------------- | -------- | ----------------- |
| Encode           | `[]byte` | `[]byte`          |
| EncodeToString   | `[]byte` | `string`          |
| EncodeFromString | `string` | `[]byte`          |
| EncodeString     | `string` | `string`          |
| Decode           | `[]byte` | `[]byte`, `error` |
| DecodeToString   | `[]byte` | `string`, `error` |
| DecodeFromString | `string` | `[]byte`, `error` |
| DecodeString     | `string` | `string`, `error` |
| IsBase32String   | `string` | `bool`            |
