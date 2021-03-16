package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type JVBStatistics struct {
	InactiveEndpoints                     int     `json:"inactive_endpoints"`
	InactiveConferences                   int     `json:"inactive_conferences"`
	TotalIceSucceededRelayed              int     `json:"total_ice_succeeded_relayed"`
	TotalLossDegradedParticipantSeconds   int     `json:"total_loss_degraded_participant_seconds"`
	BitRateDownload                       float64 `json:"bit_rate_download"`
	MucClientsConnected                   int     `json:"muc_clients_connected"`
	TotalParticipants                     int     `json:"total_participants"`
	TotalPacketsReceived                  int     `json:"total_packets_received"`
	RttAggregate                          int     `json:"rtt_aggregate"`
	PacketRateUpload                      int     `json:"packet_rate_upload"`
	P2PConferences                        int     `json:"p2p_conferences"`
	TotalLossLimitedParticipantSeconds    int     `json:"total_loss_limited_participant_seconds"`
	OctoSendBitrate                       int     `json:"octo_send_bitrate"`
	TotalDominantSpeakerChanges           int     `json:"total_dominant_speaker_changes"`
	ReceiveOnlyEndpoints                  int     `json:"receive_only_endpoints"`
	TotalColibriWebSocketMessagesReceived int     `json:"total_colibri_web_socket_messages_received"`
	OctoReceiveBitrate                    int     `json:"octo_receive_bitrate"`
	Version                               string  `json:"version"`
	TotalIceSucceeded                     int     `json:"total_ice_succeeded"`
	TotalColibriWebSocketMessagesSent     int     `json:"total_colibri_web_socket_messages_sent"`
	TotalBytesSentOcto                    int     `json:"total_bytes_sent_octo"`
	TotalDataChannelMessagesReceived      int     `json:"total_data_channel_messages_received"`
	TotalConferenceSeconds                int     `json:"total_conference_seconds"`
	NumEpsOversending                     int     `json:"num_eps_oversending"`
	BitRateUpload                         float64 `json:"bit_rate_upload"`
	TotalConferencesCompleted             int     `json:"total_conferences_completed"`
	OctoConferences                       int     `json:"octo_conferences"`
	NumEpsNoMsgTransportAfterDelay        int     `json:"num_eps_no_msg_transport_after_delay"`
	EndpointsSendingVideo                 int     `json:"endpoints_sending_video"`
	PacketRateDownload                    int     `json:"packet_rate_download"`
	MucClientsConfigured                  int     `json:"muc_clients_configured"`
	OutgoingLoss                          int     `json:"outgoing_loss"`
	OverallLoss                           int     `json:"overall_loss"`
	ConferenceSizes                       []int   `json:"conference_sizes"`
	TotalPacketsSentOcto                  int     `json:"total_packets_sent_octo"`
	ConferencesByVideoSenders             []int   `json:"conferences_by_video_senders"`
	StressLevel                           int     `json:"stress_level"`
	JitterAggregate                       int     `json:"jitter_aggregate"`
	TotalIceSucceededTCP                  int     `json:"total_ice_succeeded_tcp"`
	OctoEndpoints                         int     `json:"octo_endpoints"`
	CurrentTimestamp                      string  `json:"current_timestamp"`
	TotalPacketsDroppedOcto               int     `json:"total_packets_dropped_octo"`
	Conferences                           int     `json:"conferences"`
	Participants                          int     `json:"participants"`
	LargestConference                     int     `json:"largest_conference"`
	TotalPacketsSent                      int     `json:"total_packets_sent"`
	TotalDataChannelMessagesSent          int     `json:"total_data_channel_messages_sent"`
	IncomingLoss                          int     `json:"incoming_loss"`
	TotalBytesReceivedOcto                int     `json:"total_bytes_received_octo"`
	OctoSendPacketRate                    int     `json:"octo_send_packet_rate"`
	ConferencesByAudioSenders             []int   `json:"conferences_by_audio_senders"`
	TotalConferencesCreated               int     `json:"total_conferences_created"`
	TotalIceFailed                        int     `json:"total_ice_failed"`
	Threads                               int     `json:"threads"`
	Videochannels                         int     `json:"videochannels"`
	TotalPacketsReceivedOcto              int     `json:"total_packets_received_octo"`
	GracefulShutdown                      bool    `json:"graceful_shutdown"`
	OctoReceivePacketRate                 int     `json:"octo_receive_packet_rate"`
	TotalBytesReceived                    int     `json:"total_bytes_received"`
	TotalLossControlledParticipantSeconds int     `json:"total_loss_controlled_participant_seconds"`
	TotalPartiallyFailedConferences       int     `json:"total_partially_failed_conferences"`
	EndpointsSendingAudio                 int     `json:"endpoints_sending_audio"`
	DtlsFailedEndpoints                   int     `json:"dtls_failed_endpoints"`
	TotalBytesSent                        int     `json:"total_bytes_sent"`
	MucsConfigured                        int     `json:"mucs_configured"`
	TotalFailedConferences                int     `json:"total_failed_conferences"`
	MucsJoined                            int     `json:"mucs_joined"`
}

var (
	srv = &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d", 9001),
	}
	interrupt = make(chan os.Signal, 1) // Handle the interrupts with GO routines

)

func init() {
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://google.com")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	go func() {
		for {
			select {
			case s := <-interrupt:
				log.Printf("Signal (%d) received, stopping", s)
				time.Sleep(time.Duration(2 * time.Second)) // wait specific amount seconds to close all requests ...

				srv.Shutdown(context.Background())
			}
		}
	}()

	log.Fatal(srv.ListenAndServe(), nil)
}
