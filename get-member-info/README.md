# Get Member Info
This is a command-line Go application that will take in an API key along with any number of member IDs and will then print out their names, email addresses, and phone numbers.

## Compiling
```bash
go build get_member_info.go
```

## Running
```bash
.\get_member_info.exe <api key> <memberID> [memberID...]
```

### Example
```bash
.\get_member_info.exe mykey12345 mem-ber-id123 mem-ber-id456 mem-ber-id789
```