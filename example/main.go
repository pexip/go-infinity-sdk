package main

import (
	"context"
	"fmt"
	"log"

	infinity "github.com/pexip/go-infinity-sdk"
	"github.com/pexip/go-infinity-sdk/config"
	"github.com/pexip/go-infinity-sdk/options"
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

	// Example 2: List conferences
	fmt.Println("\n=== Configuration: Conferences ===")
	conferences, err := client.Config.ListConferences(ctx, &config.ListOptions{
		BaseListOptions: options.BaseListOptions{Limit: 5},
	})
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
		fmt.Printf("Created conference: %s (ID: %d)\n", newConf.Name, newConf.ID)

		// Clean up - delete the conference we just created
		err = client.Config.DeleteConference(ctx, newConf.ID)
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
				participant.ConferenceName)
		}
	}

	// Example 5: List worker nodes
	fmt.Println("\n=== Status: Worker Nodes ===")
	workers, err := client.Status.ListWorkers(ctx, nil)
	if err != nil {
		log.Printf("Failed to list workers: %v", err)
	} else {
		fmt.Printf("Found %d worker nodes:\n", len(workers.Objects))
		for _, worker := range workers.Objects {
			fmt.Printf("- %s (%s): %d conferences, %d participants\n",
				worker.HostName, worker.Status,
				worker.Conferences, worker.Participants)
		}
	}

	// Example 6: Command API - demonstrate participant control
	if len(participants.Objects) > 0 {
		fmt.Println("\n=== Command: Participant Control ===")
		participantUUID := participants.Objects[0].UUID

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
