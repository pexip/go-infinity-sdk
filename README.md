# Pexip Infinity Go SDK

A comprehensive Go client library for the Pexip Infinity Management API, providing easy access to all four API categories: Configuration, Status, History, and Command APIs.

## Features

- **Full API Coverage**: Complete support for all Pexip Infinity Management API endpoints
- **Type-Safe**: Strongly typed Go structs for all API requests and responses
- **Authentication**: Support for Basic Auth, Token Auth, Bearer Auth, and custom authentication
- **Retry Mechanism**: Built-in exponential backoff with jitter for reliable API calls
- **Context Support**: All operations support Go context for cancellation and timeouts
- **Comprehensive Testing**: Extensive test coverage with mocks using testify/mock
- **Flexible Configuration**: Configurable base URLs, HTTP clients, and authentication methods

## Installation

```bash
go get github.com/pexip/go-infinity-sdk
```

## Quick Start

### Basic Setup

```go
package main

import (
    "context"
    "fmt"
    "log"

    infinity "github.com/pexip/go-infinity-sdk"
)

func main() {
    // Create a new client with basic authentication and retry configuration
    client, err := infinity.New(
        infinity.WithBaseURL("https://your-pexip-server.com"),
        infinity.WithBasicAuth("admin", "your-password"),
        // Optional: customize retry behavior (default is 3 retries with exponential backoff)
        infinity.WithMaxRetries(2),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Get system status
    ctx := context.Background()
    status, err := client.Status.GetSystemStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("System Status: %s, Version: %s\n", status.Status, status.Version)
}
```

### Authentication Options

#### Basic Authentication
```go
client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithBasicAuth("admin", "password"),
)
```

#### Token Authentication
```go
client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithTokenAuth("your-api-token"),
)
```

#### Bearer Authentication
```go
import "github.com/pexip/go-infinity-sdk/auth"

client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithAuth(auth.NewBearerAuth("your-bearer-token")),
)
```

#### Custom Authentication
```go
import "github.com/pexip/go-infinity-sdk/auth"

client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithAuth(auth.NewCustomAuth("X-API-Key", "your-api-key")),
)
```

### Retry Configuration

The SDK includes a robust retry mechanism with exponential backoff to handle transient failures:

#### Default Retry Behavior
```go
// Default configuration is applied automatically
client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithBasicAuth("admin", "password"),
    // Retries up to 3 times with exponential backoff
)
```

#### Custom Retry Configuration
```go
import "time"

retryConfig := &infinity.RetryConfig{
    MaxRetries:   5,                      // Maximum number of retries
    BackoffMin:   500 * time.Millisecond, // Minimum backoff duration
    BackoffMax:   30 * time.Second,       // Maximum backoff duration
    Multiplier:   2.0,                    // Backoff multiplier
    JitterFactor: 0.1,                    // Jitter factor (0.0-1.0)
}

client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithBasicAuth("admin", "password"),
    infinity.WithRetryConfig(retryConfig),
)
```

#### Convenience Options
```go
// Disable retries completely
client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithBasicAuth("admin", "password"),
    infinity.WithNoRetries(),
)

// Set only max retries (uses default for other settings)
client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithBasicAuth("admin", "password"),
    infinity.WithMaxRetries(1),
)
```

The retry mechanism automatically retries on:
- **HTTP Status Codes**: 429, 500, 502, 503, 504
- **Network Errors**: Connection refused, timeouts, DNS failures

It will **NOT** retry on:
- **Client Errors**: 4xx status codes (400, 401, 403, 404, etc.)
- **Context Cancellation**: Respects context timeouts and cancellation

## API Examples

### Configuration API

#### Managing Conferences

```go
package main

import (
    "context"
    "fmt"
    "log"

    infinity "github.com/pexip/go-infinity-sdk"
    "github.com/pexip/go-infinity-sdk/config"
)

func main() {
    client, err := infinity.New(
        infinity.WithBaseURL("https://your-pexip-server.com"),
        infinity.WithBasicAuth("admin", "password"),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Create a new conference
    createReq := &config.ConferenceCreateRequest{
        Name:        "My Test Conference",
        Description: "A test conference for demonstration",
        ServiceType: "conference",
        AllowGuests: true,
        GuestsMuted: false,
    }

    conference, err := client.Config.CreateConference(ctx, createReq)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created conference: %s (ID: %d)\n", conference.Name, conference.ID)

    // List conferences with pagination and search
    listOpts := &config.ListOptions{}
    listOpts.Limit = 10
    listOpts.Offset = 0
    listOpts.Search = "test"

    conferences, err := client.Config.ListConferences(ctx, listOpts)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d conferences\n", len(conferences.Objects))
    for _, conf := range conferences.Objects {
        fmt.Printf("- %s (ID: %d)\n", conf.Name, conf.ID)
    }

    // Update a conference
    updateReq := &config.ConferenceUpdateRequest{
        Description: "Updated description",
    }

    updatedConf, err := client.Config.UpdateConference(ctx, conference.ID, updateReq)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Updated conference description: %s\n", updatedConf.Description)

    // Delete the conference
    err = client.Config.DeleteConference(ctx, conference.ID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Conference deleted successfully")
}
```

#### Managing Locations

```go
// Create a location
createLocReq := &config.LocationCreateRequest{
    Name:        "New York Office",
    Description: "Main office location",
}

location, err := client.Config.CreateLocation(ctx, createLocReq)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Created location: %s (ID: %d)\n", location.Name, location.ID)

// List all locations
locations, err := client.Config.ListLocations(ctx, nil)
if err != nil {
    log.Fatal(err)
}

for _, loc := range locations.Objects {
    fmt.Printf("Location: %s - %s\n", loc.Name, loc.Description)
}
```

### Status API

#### System and Conference Status

```go
package main

import (
    "context"
    "fmt"
    "log"

    infinity "github.com/pexip/go-infinity-sdk"
    "github.com/pexip/go-infinity-sdk/status"
)

func main() {
    client, err := infinity.New(
        infinity.WithBaseURL("https://your-pexip-server.com"),
        infinity.WithBasicAuth("admin", "password"),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Get system status
    sysStatus, err := client.Status.GetSystemStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("System Status: %s\n", sysStatus.Status)
    fmt.Printf("Version: %s\n", sysStatus.Version)
    fmt.Printf("Uptime: %d seconds\n", sysStatus.Uptime)
    fmt.Printf("CPU Load: %.2f%%\n", sysStatus.CPULoad)
    fmt.Printf("Memory Usage: %d/%d MB\n", 
        sysStatus.UsedMemory/1024/1024, 
        sysStatus.TotalMemory/1024/1024)

    // List active conferences
    conferences, err := client.Status.ListConferences(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\nActive Conferences: %d\n", len(conferences.Objects))
    for _, conf := range conferences.Objects {
        fmt.Printf("- %s: %d participants (Started: %t)\n", 
            conf.Name, conf.ParticipantCount, conf.Started)
    }

    // List participants with pagination
    participantOpts := &status.ListOptions{}
    participantOpts.Limit = 20
    participants, err := client.Status.ListParticipants(ctx, participantOpts)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\nActive Participants: %d\n", len(participants.Objects))
    for _, participant := range participants.Objects {
        fmt.Printf("- %s (%s) in %s - Muted: %t\n", 
            participant.DisplayName, 
            participant.Role,
            participant.ConferenceName,
            participant.IsMuted)
    }

    // List worker nodes
    workers, err := client.Status.ListWorkers(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\nWorker Nodes: %d\n", len(workers.Objects))
    for _, worker := range workers.Objects {
        fmt.Printf("- %s (%s): CPU %.1f%%, Memory %.1f%%, %d conferences\n", 
            worker.HostName, 
            worker.Status,
            worker.CPU,
            worker.Memory,
            worker.Conferences)
    }

    // List system alarms
    alarms, err := client.Status.ListAlarms(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\nSystem Alarms: %d\n", len(alarms.Objects))
    for _, alarm := range alarms.Objects {
        status := "ACTIVE"
        if alarm.Cleared {
            status = "CLEARED"
        }
        fmt.Printf("- [%s] %s: %s (%s)\n", 
            alarm.Level, 
            alarm.Name, 
            alarm.Details,
            status)
    }
}
```

### History API

#### Conference and Participant History

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    infinity "github.com/pexip/go-infinity-sdk"
    "github.com/pexip/go-infinity-sdk/history"
)

func main() {
    client, err := infinity.New(
        infinity.WithBaseURL("https://your-pexip-server.com"),
        infinity.WithBasicAuth("admin", "password"),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Get conference history for the last 24 hours
    yesterday := time.Now().Add(-24 * time.Hour)
    listOpts := &history.ListOptions{}
    listOpts.Limit = 50
    listOpts.StartTime = &yesterday
    listOpts.Search = "weekly"

    conferences, err := client.History.ListConferences(ctx, listOpts)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Conference History (last 24h): %d conferences\n", len(conferences.Objects))
    for _, conf := range conferences.Objects {
        duration := time.Duration(conf.DurationSeconds) * time.Second
        fmt.Printf("- %s: %v duration, %d participants\n", 
            conf.Name, 
            duration,
            conf.TotalParticipants)
    }

    // Get participant history for a specific conference
    if len(conferences.Objects) > 0 {
        confID := conferences.Objects[0].ID
        participants, err := client.History.ListParticipantsByConference(ctx, confID, nil)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("\nParticipants in conference %d:\n", confID)
        for _, participant := range participants.Objects {
            duration := time.Duration(participant.DurationSeconds) * time.Second
            fmt.Printf("- %s (%s): %v duration, %s reason\n",
                participant.DisplayName,
                participant.Role,
                duration,
                participant.DisconnectReason)
        }

        // Get media stream details for the first participant
        if len(participants.Objects) > 0 {
            participantID := participants.Objects[0].ID
            streams, err := client.History.ListMediaStreamsByParticipant(ctx, participantID, nil)
            if err != nil {
                log.Fatal(err)
            }

            fmt.Printf("\nMedia streams for participant %d:\n", participantID)
            for _, stream := range streams.Objects {
                fmt.Printf("- %s %s: %s codec, %d kbps\n",
                    stream.StreamType,
                    stream.Direction,
                    stream.Codec,
                    stream.Bitrate/1000)
            }
        }
    }
}
```

### Command API

#### Conference and Participant Control

```go
package main

import (
    "context"
    "fmt"
    "log"

    infinity "github.com/pexip/go-infinity-sdk"
    "github.com/pexip/go-infinity-sdk/command"
)

func main() {
    client, err := infinity.New(
        infinity.WithBaseURL("https://your-pexip-server.com"),
        infinity.WithBasicAuth("admin", "password"),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Get current participants to demonstrate commands
    participants, err := client.Status.ListParticipants(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    if len(participants.Objects) == 0 {
        fmt.Println("No active participants found")
        return
    }

    participant := participants.Objects[0]
    participantUUID := participant.UUID
    conferenceID := participant.ConferenceID

    fmt.Printf("Managing participant: %s (UUID: %s)\n", participant.DisplayName, participantUUID)

    // Mute a participant
    result, err := client.Command.MuteParticipant(ctx, participantUUID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Mute result: %s - %s\n", result.Status, result.Message)

    // Send a message to the participant
    result, err = client.Command.SendMessageToParticipant(ctx, participantUUID, "Welcome to the conference!")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Message sent: %s\n", result.Status)

    // Spotlight the participant
    result, err = client.Command.SpotlightParticipant(ctx, participantUUID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Spotlight result: %s\n", result.Status)

    // Promote participant to chair
    result, err = client.Command.PromoteParticipant(ctx, participantUUID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Promotion result: %s\n", result.Status)

    // Lock the conference
    result, err = client.Command.LockConference(ctx, conferenceID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Conference lock result: %s\n", result.Status)

    // Send message to all participants in the conference
    result, err = client.Command.SendMessageToConference(ctx, conferenceID, "This conference is now locked")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Conference message result: %s\n", result.Status)

    // Transfer participant to another conference
    transferOpts := &command.TransferOptions{
        Role: "guest",
    }
    result, err = client.Command.TransferParticipant(ctx, participantUUID, "other-conference", transferOpts)
    if err != nil {
        log.Printf("Transfer failed (expected if 'other-conference' doesn't exist): %v\n", err)
    } else {
        fmt.Printf("Transfer result: %s\n", result.Status)
    }

    // Start a conference programmatically
    result, err = client.Command.StartConference(ctx, "scheduled-meeting")
    if err != nil {
        log.Printf("Start conference failed (expected if conference doesn't exist): %v\n", err)
    } else {
        fmt.Printf("Conference start result: %s\n", result.Status)
    }
}
```

## Advanced Usage

### Custom HTTP Client

```go
import (
    "net/http"
    "time"
)

// Create a custom HTTP client with timeout and custom transport
httpClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}

client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithHTTPClient(httpClient),
    infinity.WithBasicAuth("admin", "password"),
)
```

### Custom HTTP Transport

For more fine-grained control over HTTP transport settings (proxies, TLS configuration, connection pooling), you can use `WithTransport`:

```go
import (
    "crypto/tls"
    "net/http"
    "net/url"
    "time"
)

// Example 1: Custom transport with proxy and TLS settings
proxyURL, _ := url.Parse("http://proxy.company.com:8080")
customTransport := &http.Transport{
    Proxy: http.ProxyURL(proxyURL),
    TLSClientConfig: &tls.Config{
        InsecureSkipVerify: true, // Only for testing!
    },
    MaxIdleConns:        50,
    MaxIdleConnsPerHost: 10,
    IdleConnTimeout:     30 * time.Second,
    DisableCompression:  true,
}

client, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithTransport(customTransport),
    infinity.WithBasicAuth("admin", "password"),
)

// Example 2: Transport with client certificates
cert, _ := tls.LoadX509KeyPair("client.crt", "client.key")
tlsTransport := &http.Transport{
    TLSClientConfig: &tls.Config{
        Certificates: []tls.Certificate{cert},
    },
}

tlsClient, err := infinity.New(
    infinity.WithBaseURL("https://your-pexip-server.com"),
    infinity.WithTransport(tlsTransport),
    infinity.WithBasicAuth("admin", "password"),
)
```

### Error Handling

```go
import "github.com/pexip/go-infinity-sdk"

// All API calls return typed errors
conferences, err := client.Config.ListConferences(ctx, nil)
if err != nil {
    // Check if it's an API error
    if apiErr, ok := err.(*infinity.APIError); ok {
        fmt.Printf("API Error %d: %s\n", apiErr.StatusCode, apiErr.Message)
        if apiErr.Details != "" {
            fmt.Printf("Details: %s\n", apiErr.Details)
        }
    } else {
        // Handle other types of errors (network, timeout, etc.)
        fmt.Printf("Other error: %v\n", err)
    }
}
```

### Context and Timeouts

```go
import (
    "context"
    "time"
)

// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// Use the context for all API calls
conferences, err := client.Config.ListConferences(ctx, nil)
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println("Request timed out")
    }
}
```

## API Reference

### Configuration API (`client.Config`)

- **Conferences**
  - `ListConferences(ctx, opts)` - List conferences
  - `GetConference(ctx, id)` - Get conference by ID
  - `CreateConference(ctx, req)` - Create new conference
  - `UpdateConference(ctx, id, req)` - Update conference
  - `DeleteConference(ctx, id)` - Delete conference

- **Locations**
  - `ListLocations(ctx, opts)` - List locations
  - `GetLocation(ctx, id)` - Get location by ID
  - `CreateLocation(ctx, req)` - Create new location
  - `UpdateLocation(ctx, id, req)` - Update location
  - `DeleteLocation(ctx, id)` - Delete location

### Status API (`client.Status`)

- `GetSystemStatus(ctx)` - Get overall system status
- `ListConferences(ctx, opts)` - List active conferences
- `GetConference(ctx, id)` - Get conference status by ID
- `ListParticipants(ctx, opts)` - List active participants
- `GetParticipant(ctx, uuid)` - Get participant by UUID
- `ListWorkers(ctx, opts)` - List worker nodes
- `GetWorker(ctx, nodeID)` - Get worker by node ID
- `ListAlarms(ctx, opts)` - List system alarms
- `GetAlarm(ctx, id)` - Get alarm by ID

### History API (`client.History`)

- **Conferences**
  - `ListConferences(ctx, opts)` - List conference history
  - `GetConference(ctx, id)` - Get conference history by ID

- **Participants**
  - `ListParticipants(ctx, opts)` - List participant history
  - `GetParticipant(ctx, id)` - Get participant history by ID
  - `ListParticipantsByConference(ctx, confID, opts)` - Get participants for a conference

- **Media Streams**
  - `ListMediaStreams(ctx, opts)` - List media stream history
  - `GetMediaStream(ctx, id)` - Get media stream by ID
  - `ListMediaStreamsByParticipant(ctx, participantID, opts)` - Get streams for a participant

### Command API (`client.Command`)

- **Participant Control**
  - `DisconnectParticipant(ctx, uuid)` - Disconnect participant
  - `MuteParticipant(ctx, uuid)` - Mute participant
  - `UnmuteParticipant(ctx, uuid)` - Unmute participant
  - `ToggleMuteParticipant(ctx, uuid)` - Toggle mute status
  - `SpotlightParticipant(ctx, uuid)` - Enable spotlight
  - `UnspotlightParticipant(ctx, uuid)` - Disable spotlight
  - `ToggleSpotlightParticipant(ctx, uuid)` - Toggle spotlight
  - `PromoteParticipant(ctx, uuid)` - Promote to chair
  - `DemoteParticipant(ctx, uuid)` - Demote to guest
  - `ChangeParticipantRole(ctx, uuid, role)` - Change role
  - `SendMessageToParticipant(ctx, uuid, message)` - Send message
  - `TransferParticipant(ctx, uuid, conference, opts)` - Transfer participant

- **Conference Control**
  - `LockConference(ctx, id)` - Lock conference
  - `UnlockConference(ctx, id)` - Unlock conference
  - `ToggleLockConference(ctx, id)` - Toggle lock status
  - `SendMessageToConference(ctx, id, message)` - Send message to all
  - `StartConference(ctx, alias)` - Start conference
  - `StopConference(ctx, id)` - Stop conference

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass: `go test ./...`
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For questions and support, please refer to the [Pexip Infinity API Documentation](https://docs.pexip.com/api_manage/) or open an issue in this repository.