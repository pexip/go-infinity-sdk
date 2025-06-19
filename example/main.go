package main

import (
	"context"
	"fmt"
	"log"

	infinity "github.com/pexip/go-infinity-sdk/v38"
	"github.com/pexip/go-infinity-sdk/v38/config"
)

func main() {
	// Create a new Infinity API client
	client, err := infinity.New(
		infinity.WithBaseURL("https://your-pexip-server.com"),
		infinity.WithBasicAuth("admin", "your-password"),
	)
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}

	ctx := context.Background()

	// Example 1: Get system status
	fmt.Println("=== System Status ===")
	status, err := client.Status.GetSystemStatus(ctx)
	if err != nil {
		log.Printf("Failed to get system status: %v", err)
	} else {
		fmt.Printf("Status: %s, Version: %s, Uptime: %d seconds\n",
			status.Status, status.Version, status.Uptime)
	}

	// Example 2: List conferences with search and pagination
	fmt.Println("\n=== Configuration: Conferences ===")
	listOpts := &config.ListOptions{}
	listOpts.Limit = 5
	listOpts.Search = "example" // Optional: search for conferences containing "example"
	conferences, err := client.Config.ListConferences(ctx, listOpts)
	if err != nil {
		log.Printf("Failed to list conferences: %v", err)
	} else {
		fmt.Printf("Found %d conferences:\n", len(conferences.Objects))
		for _, conf := range conferences.Objects {
			fmt.Printf("- %s (ID: %d)\n", conf.Name, conf.ID)
		}
	}

	// Example 3: Create a new conference
	fmt.Println("\n=== Creating a Conference ===")
	createReq := &config.ConferenceCreateRequest{
		Name:        "Go SDK Example Conference",
		Description: "Conference created by Go SDK example",
		ServiceType: "conference",
		AllowGuests: true,
		GuestsMuted: false,
	}

	newConf, err := client.Config.CreateConference(ctx, createReq)
	if err != nil {
		log.Printf("Failed to create conference: %v", err)
	} else {
		id, _ := newConf.ResourceID()
		fmt.Printf("Created conference (ID: %d)\n", id)

		// Clean up - delete the conference we just created
		err = client.Config.DeleteConference(ctx, id)
		if err != nil {
			log.Printf("Failed to delete conference: %v", err)
		} else {
			fmt.Println("Conference cleaned up successfully")
		}
	}

	// Example 4: List active participants
	fmt.Println("\n=== Status: Active Participants ===")
	participants, err := client.Status.ListParticipants(ctx, nil)
	if err != nil {
		log.Printf("Failed to list participants: %v", err)
	} else {
		fmt.Printf("Found %d active participants:\n", len(participants.Objects))
		for _, participant := range participants.Objects {
			fmt.Printf("- %s (%s) in %s\n",
				participant.DisplayName,
				participant.Role,
				participant.Conference)
		}
	}

	// Example 5: List worker VMs
	fmt.Println("\n=== Status: Worker VMs ===")
	workers, err := client.Status.ListWorkerVMs(ctx, nil)
	if err != nil {
		log.Printf("Failed to list worker VMs: %v", err)
	} else {
		fmt.Printf("Found %d worker VMs:\n", len(workers.Objects))
		for _, worker := range workers.Objects {
			fmt.Printf("- %s (%s): Media load: %d, Signaling count: %d\n",
				worker.Name, worker.SyncStatus,
				worker.MediaLoad, worker.SignalingCount)
		}
	}

	// Example 6: Command API - demonstrate participant control
	if len(participants.Objects) > 0 {
		fmt.Println("\n=== Command: Participant Control ===")
		participantUUID := participants.Objects[0].CallUUID

		// Send a welcome message
		result, err := client.Command.SendMessageToParticipant(ctx, participantUUID, "Welcome! This message was sent via the Go SDK.")
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			fmt.Printf("Message sent successfully: %s\n", result.Status)
		}
	}

	fmt.Println("\n=== Example completed ===")
}
