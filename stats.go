package main

type JVBStatistics struct {
	InactiveEndpoints                     int     `json:"inactive_endpoints" prom:"Gauge;Number of inactive endpoints"`
	InactiveConferences                   int     `json:"inactive_conferences" prom:"Gauge;The number of inactive conferences (with no endpoints sending audio or video)"`
	TotalIceSucceededRelayed              int     `json:"total_ice_succeeded_relayed" prom:"Counter;Total number of times ICE succeeded."`
	TotalLossDegradedParticipantSeconds   int     `json:"total_loss_degraded_participant_seconds" prom:"Counter;The total number of participant-seconds that are loss-degraded."`
	BitRateDownload                       float64 `json:"bit_rate_download" prom:"Gauge;The total incoming bitrate for the video bridge in kilobits per second."`
	MucClientsConnected                   int     `json:"muc_clients_connected" prom:"Gauge;Number of configured MUC clients that are connected to XMPP."`
	TotalParticipants                     int     `json:"total_participants" prom:"Gauge;The current number of participants."`
	TotalPacketsReceived                  int     `json:"total_packets_received" prom:"Counter;Total number of packets received"`
	RttAggregate                          float64 `json:"rtt_aggregate" prom:"Gauge;An average value (in milliseconds) of the RTT across all streams."`
	PacketRateUpload                      float64 `json:"packet_rate_upload" prom:"Gauge;The total outgoing packet rate for the video bridge in packets per second."`
	P2PConferences                        int     `json:"p2p_conferences" prom:"Gauge;The current number of p2p conferences."`
	TotalLossLimitedParticipantSeconds    int     `json:"total_loss_limited_participant_seconds" prom:"Counter;The total number of participant-seconds that are loss-limited."`
	OctoSendBitrate                       int     `json:"octo_send_bitrate" prom:"Gauge;The total outgoing bitrate for the OCTO video bridge in kilobits per second."`
	TotalDominantSpeakerChanges           int     `json:"total_dominant_speaker_changes" prom:"Counter;The cumulative number of times the dominant speaker changed."`
	ReceiveOnlyEndpoints                  int     `json:"receive_only_endpoints" prom:"Gauge;The number of endpoints endpoints that have no audio or video, but are not inactive"`
	TotalColibriWebSocketMessagesReceived int     `json:"total_colibri_web_socket_messages_received" prom:"Counter;The total number messages received through COLIBRI web sockets."`
	OctoReceiveBitrate                    int     `json:"octo_receive_bitrate" prom:"Gauge;The total receiving bitrate for the OCTO video bridge in kilobits per second."`
	Version                               string  `json:"version"`
	TotalIceSucceeded                     int     `json:"total_ice_succeeded" prom:"Counter;Total number of times ICE succeeded."`
	TotalColibriWebSocketMessagesSent     int     `json:"total_colibri_web_socket_messages_sent" prom:"Counter;The total number messages sent through COLIBRI web sockets."`
	TotalBytesSentOcto                    int     `json:"total_bytes_sent_octo" prom:"Gauge;The total outgoing bit rate for the OCTO video bridge in bytes per second."`
	TotalDataChannelMessagesReceived      int     `json:"total_data_channel_messages_received" prom:"Counter;The total number messages received through data channels."`
	TotalConferenceSeconds                int     `json:"total_conference_seconds" prom:"Counter;The sum of the lengths of all completed conferences, in seconds."`
	NumEpsOversending                     int     `json:"num_eps_oversending"`
	BitRateUpload                         float64 `json:"bit_rate_upload" prom:"Gauge;The total outgoing bitrate for the video bridge in kilobits per second."`
	TotalConferencesCompleted             int     `json:"total_conferences_completed" prom:"Counter;The total number of conferences completed on the bridge"`
	OctoConferences                       int     `json:"octo_conferences" prom:"Gauge;The current number of OCTO conferences."`
	NumEpsNoMsgTransportAfterDelay        int     `json:"num_eps_no_msg_transport_after_delay"`
	EndpointsSendingVideo                 int     `json:"endpoints_sending_video" prom:"Gauge;Endpoint sending video"`
	PacketRateDownload                    float64 `json:"packet_rate_download" prom:"Gauge;The total incoming packet rate for the video bridge in packets per second."`
	MucClientsConfigured                  int     `json:"muc_clients_configured" prom:"Gauge;Number of configured MUC clients."`
	OutgoingLoss                          float64 `json:"outgoing_loss" prom:"Gauge;The fraction of lost outgoing RTP packets. This is based on incoming RTCP Receiver Reports, and an attempt to subtract the fraction of packets that were not sent (i.e. were lost before they reached the bridge). Further, this is averaged over all streams of all users as opposed to all packets, so it is not correctly weighted. This is not accurate, but may be a useful metric nonetheless."`
	OverallLoss                           float64 `json:"overall_loss"`
	ConferenceSizes                       []int   `json:"conference_sizes" prom:"Untyped;Histogram of conference sizes"`
	TotalPacketsSentOcto                  int     `json:"total_packets_sent_octo" prom:"Gauge;The total of sent dropped packets handled by the OCTO video bridge."`
	ConferencesByVideoSenders             []int   `json:"conferences_by_video_senders" prom:"Untyped;Histogram of conference by video senders"`
	StressLevel                           float64 `json:"stress_level" prom:"Gauge;Stress Level reported to Jicofo by the videobridge."`
	JitterAggregate                       float64 `json:"jitter_aggregate" prom:"Gauge;Experimental. An average value (in milliseconds) of the jitter calculated for incoming and outgoing streams. This hasn't been tested and it is currently not known whether the values are correct or not."`
	TotalIceSucceededTCP                  int     `json:"total_ice_succeeded_tcp" prom:"Counter; Total number of times ICE succeeded over TCP."`
	OctoEndpoints                         int     `json:"octo_endpoints" prom:"Gauge;The current number of OCTO endpoints."`
	CurrentTimestamp                      string  `json:"current_timestamp"`
	TotalPacketsDroppedOcto               int     `json:"total_packets_dropped_octo" prom:"Gauge;The total of dropped packets handled by the OCTO video bridge."`
	Conferences                           int     `json:"conferences" prom:"Gauge;The current number of conferences."`
	Participants                          int     `json:"participants" prom:"Gauge;The current number of participants."`
	LargestConference                     int     `json:"largest_conference" prom:"Gauge;The number of participants in the largest conference currently hosted on the bridge."`
	TotalPacketsSent                      int     `json:"total_packets_sent" prom:"Gauge;The total of sent packets handled by video bridge."`
	TotalDataChannelMessagesSent          int     `json:"total_data_channel_messages_sent" prom:"Counter;The total number messages sent through data channels."`
	IncomingLoss                          float64 `json:"incoming_loss" prom:"Gauge;The fraction of lost incoming RTP packets. This is based on RTP sequence numbers and is relatively accurate."`
	TotalBytesReceivedOcto                int     `json:"total_bytes_received_octo" prom:"Gauge;The total incoming bit rate for the OCTO video bridge in bytes per second."`
	OctoSendPacketRate                    int     `json:"octo_send_packet_rate" prom:"Gauge;The total outgoing packet rate for the OCTO video bridge in packets per second."`
	ConferencesByAudioSenders             []int   `json:"conferences_by_audio_senders" prom:"Untyped;Histogram of conference by audio senders"`
	TotalConferencesCreated               int     `json:"total_conferences_created" prom:"Counter;The total number of conferences created on the bridge."`
	TotalIceFailed                        int     `json:"total_ice_failed" prom:"Counter;total number of times ICE failed."`
	Threads                               int     `json:"threads" prom:"Gauge;The number of Java threads that the video bridge is using."`
	Videochannels                         int     `json:"videochannels" prom:"Gauge;The current number of video channels."`
	Audiochannels                         int     `json:"audiochannels" prom:"Gauge;The current number of audio channels."`
	TotalPacketsReceivedOcto              int     `json:"total_packets_received_octo" prom:"Gauge;The total of incoming dropped packets handled by the OCTO video bridge."`
	GracefulShutdown                      bool    `json:"graceful_shutdown"` //TODO: Find out more regarding this metrics
	OctoReceivePacketRate                 int     `json:"octo_receive_packet_rate" prom:"Gauge;The total incoming packet rate for the OCTO video bridge in packets per second."`
	TotalBytesReceived                    int     `json:"total_bytes_received" prom:"Counter;Total bytes received."`
	TotalLossControlledParticipantSeconds int     `json:"total_loss_controlled_participant_seconds" prom:"Counter;The total number of participant-seconds that are loss-controlled."`
	TotalPartiallyFailedConferences       int     `json:"total_partially_failed_conferences" prom:"Counter;The total number of partially failed conferences on the bridge. A conference is marked as partially failed when some of its channels has failed. A channel is marked as failed if it had no payload activity."`
	EndpointsSendingAudio                 int     `json:"endpoints_sending_audio" prom:"Gauge;Endpoint sending audio"`
	DtlsFailedEndpoints                   int     `json:"dtls_failed_endpoints"` //TODO: Find out more regarding this metrics
	TotalBytesSent                        int     `json:"total_bytes_sent" prom:"Counter;The number of total bytes sent."`
	MucsConfigured                        int     `json:"mucs_configured" prom:"Gauge;Number of MUCs that are configured"`
	TotalFailedConferences                int     `json:"total_failed_conferences" prom:"Counter;The total number of failed conferences on the bridge. A conference is marked as failed when all of its channels have failed. A channel is marked as failed if it had no payload activity."`
	MucsJoined                            int     `json:"mucs_joined" prom:"Gauge;Number of MUCs that are joined."`
}
