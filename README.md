# jitsi-jvb-prometheus-adapter

## Motivation
The jitsi-jvb-prometheus-adapter has been created to:
* Export metrics using Prometheus client
* Leverage dynamic code to collect and deliver metrics
* Support histograms for metrics from Jitsi JVB

## Running
> TESTED WITH "stable-5390-3"

The designed way to use jitsi-jvb-prometheus-adapter is to run it side by side JVB , so in kubernetes that would be as one of the containers in the pod.
```yaml
      containers:
        - name: jvb
          image: jitsi/jvb:stable-5390-3
          # ...... removed for visibility
        - name: prometheus-exporter
          image: rafpe/jitsi-jvb-prometheus-adapter:latest # in production recommended to pin to specific version
          imagePullPolicy: IfNotPresent
          ports:
          - containerPort: 9001
            name: metrics
          env:
            - name: JVB_STATS_URL
              value: "http://localhost:8080/colibri/stats" # change accordingly to your JVB setup ( IP and PORT )

```

## Exposed metrics
The following are metric exposed by exporter

```
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 5.96e-05
go_gc_duration_seconds{quantile="0.25"} 0.0002589
go_gc_duration_seconds{quantile="0.5"} 0.0004578
go_gc_duration_seconds{quantile="0.75"} 0.0007894
go_gc_duration_seconds{quantile="1"} 0.0313565
go_gc_duration_seconds_sum 0.1930966
go_gc_duration_seconds_count 287
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 25
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.15"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 5.1296e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 9.96735096e+08
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.503094e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 8.454455e+06
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 2.3355096571563482e-05
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 5.102544e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 5.1296e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 5.988352e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 6.504448e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 11643
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 5.7704448e+07
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.6387968e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.6166145957113123e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 8.466098e+06
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 6944
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 85272
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 114688
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 8.269232e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 948410
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 720896
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 720896
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.4793984e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 12
# HELP jitsi_audiochannels The current number of audio channels.
# TYPE jitsi_audiochannels gauge
jitsi_audiochannels 0
# HELP jitsi_bit_rate_download The total incoming bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_download gauge
jitsi_bit_rate_download 0.5
# HELP jitsi_bit_rate_upload The total outgoing bitrate for the video bridge in kilobits per second.
# TYPE jitsi_bit_rate_upload gauge
jitsi_bit_rate_upload 0.5
# HELP jitsi_conference_sizes Histogram of conference sizes
# TYPE jitsi_conference_sizes histogram
jitsi_conference_sizes_bucket{le="0"} 0
jitsi_conference_sizes_bucket{le="1"} 0
jitsi_conference_sizes_bucket{le="2"} 0
jitsi_conference_sizes_bucket{le="3"} 0
jitsi_conference_sizes_bucket{le="4"} 0
jitsi_conference_sizes_bucket{le="5"} 0
jitsi_conference_sizes_bucket{le="6"} 0
jitsi_conference_sizes_bucket{le="7"} 0
jitsi_conference_sizes_bucket{le="8"} 0
jitsi_conference_sizes_bucket{le="9"} 0
jitsi_conference_sizes_bucket{le="10"} 0
jitsi_conference_sizes_bucket{le="11"} 0
jitsi_conference_sizes_bucket{le="12"} 0
jitsi_conference_sizes_bucket{le="13"} 0
jitsi_conference_sizes_bucket{le="14"} 0
jitsi_conference_sizes_bucket{le="15"} 0
jitsi_conference_sizes_bucket{le="16"} 0
jitsi_conference_sizes_bucket{le="17"} 0
jitsi_conference_sizes_bucket{le="18"} 0
jitsi_conference_sizes_bucket{le="19"} 0
jitsi_conference_sizes_bucket{le="20"} 0
jitsi_conference_sizes_bucket{le="+Inf"} 0
jitsi_conference_sizes_sum 0
jitsi_conference_sizes_count 0
# HELP jitsi_conferences The current number of conferences.
# TYPE jitsi_conferences gauge
jitsi_conferences 0
# HELP jitsi_conferences_by_audio_senders Histogram of conference by audio senders
# TYPE jitsi_conferences_by_audio_senders histogram
jitsi_conferences_by_audio_senders_bucket{le="0"} 0
jitsi_conferences_by_audio_senders_bucket{le="1"} 0
jitsi_conferences_by_audio_senders_bucket{le="2"} 0
jitsi_conferences_by_audio_senders_bucket{le="3"} 0
jitsi_conferences_by_audio_senders_bucket{le="4"} 0
jitsi_conferences_by_audio_senders_bucket{le="5"} 0
jitsi_conferences_by_audio_senders_bucket{le="6"} 0
jitsi_conferences_by_audio_senders_bucket{le="7"} 0
jitsi_conferences_by_audio_senders_bucket{le="8"} 0
jitsi_conferences_by_audio_senders_bucket{le="9"} 0
jitsi_conferences_by_audio_senders_bucket{le="10"} 0
jitsi_conferences_by_audio_senders_bucket{le="11"} 0
jitsi_conferences_by_audio_senders_bucket{le="12"} 0
jitsi_conferences_by_audio_senders_bucket{le="13"} 0
jitsi_conferences_by_audio_senders_bucket{le="14"} 0
jitsi_conferences_by_audio_senders_bucket{le="15"} 0
jitsi_conferences_by_audio_senders_bucket{le="16"} 0
jitsi_conferences_by_audio_senders_bucket{le="17"} 0
jitsi_conferences_by_audio_senders_bucket{le="18"} 0
jitsi_conferences_by_audio_senders_bucket{le="19"} 0
jitsi_conferences_by_audio_senders_bucket{le="20"} 0
jitsi_conferences_by_audio_senders_bucket{le="+Inf"} 0
jitsi_conferences_by_audio_senders_sum 0
jitsi_conferences_by_audio_senders_count 0
# HELP jitsi_conferences_by_video_senders Histogram of conference by video senders
# TYPE jitsi_conferences_by_video_senders histogram
jitsi_conferences_by_video_senders_bucket{le="0"} 0
jitsi_conferences_by_video_senders_bucket{le="1"} 0
jitsi_conferences_by_video_senders_bucket{le="2"} 0
jitsi_conferences_by_video_senders_bucket{le="3"} 0
jitsi_conferences_by_video_senders_bucket{le="4"} 0
jitsi_conferences_by_video_senders_bucket{le="5"} 0
jitsi_conferences_by_video_senders_bucket{le="6"} 0
jitsi_conferences_by_video_senders_bucket{le="7"} 0
jitsi_conferences_by_video_senders_bucket{le="8"} 0
jitsi_conferences_by_video_senders_bucket{le="9"} 0
jitsi_conferences_by_video_senders_bucket{le="10"} 0
jitsi_conferences_by_video_senders_bucket{le="11"} 0
jitsi_conferences_by_video_senders_bucket{le="12"} 0
jitsi_conferences_by_video_senders_bucket{le="13"} 0
jitsi_conferences_by_video_senders_bucket{le="14"} 0
jitsi_conferences_by_video_senders_bucket{le="15"} 0
jitsi_conferences_by_video_senders_bucket{le="16"} 0
jitsi_conferences_by_video_senders_bucket{le="17"} 0
jitsi_conferences_by_video_senders_bucket{le="18"} 0
jitsi_conferences_by_video_senders_bucket{le="19"} 0
jitsi_conferences_by_video_senders_bucket{le="20"} 0
jitsi_conferences_by_video_senders_bucket{le="+Inf"} 0
jitsi_conferences_by_video_senders_sum 0
jitsi_conferences_by_video_senders_count 0
# HELP jitsi_endpoints_sending_audio Endpoint sending audio
# TYPE jitsi_endpoints_sending_audio gauge
jitsi_endpoints_sending_audio 0
# HELP jitsi_endpoints_sending_video Endpoint sending video
# TYPE jitsi_endpoints_sending_video gauge
jitsi_endpoints_sending_video 0
# HELP jitsi_inactive_conferences The number of inactive conferences (with no endpoints sending audio or video)
# TYPE jitsi_inactive_conferences gauge
jitsi_inactive_conferences 0
# HELP jitsi_inactive_endpoints Number of inactive endpoints
# TYPE jitsi_inactive_endpoints gauge
jitsi_inactive_endpoints 0
# HELP jitsi_incoming_loss The fraction of lost incoming RTP packets. This is based on RTP sequence numbers and is relatively accurate.
# TYPE jitsi_incoming_loss gauge
jitsi_incoming_loss 0
# HELP jitsi_jitter_aggregate Experimental. An average value (in milliseconds) of the jitter calculated for incoming and outgoing streams. This hasn't been tested and it is currently not known whether the values are correct or not.
# TYPE jitsi_jitter_aggregate gauge
jitsi_jitter_aggregate 0
# HELP jitsi_largest_conference The number of participants in the largest conference currently hosted on the bridge.
# TYPE jitsi_largest_conference gauge
jitsi_largest_conference 0
# HELP jitsi_muc_clients_configured Number of configured MUC clients.
# TYPE jitsi_muc_clients_configured gauge
jitsi_muc_clients_configured 1
# HELP jitsi_muc_clients_connected Number of configured MUC clients that are connected to XMPP.
# TYPE jitsi_muc_clients_connected gauge
jitsi_muc_clients_connected 1
# HELP jitsi_mucs_configured Number of MUCs that are configured
# TYPE jitsi_mucs_configured gauge
jitsi_mucs_configured 1
# HELP jitsi_mucs_joined Number of MUCs that are joined.
# TYPE jitsi_mucs_joined gauge
jitsi_mucs_joined 1
# HELP jitsi_octo_conferences The current number of OCTO conferences.
# TYPE jitsi_octo_conferences gauge
jitsi_octo_conferences 0
# HELP jitsi_octo_endpoints The current number of OCTO endpoints.
# TYPE jitsi_octo_endpoints gauge
jitsi_octo_endpoints 0
# HELP jitsi_octo_receive_bitrate The total receiving bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_receive_bitrate gauge
jitsi_octo_receive_bitrate 0
# HELP jitsi_octo_receive_packet_rate The total incoming packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_receive_packet_rate gauge
jitsi_octo_receive_packet_rate 0
# HELP jitsi_octo_send_bitrate The total outgoing bitrate for the OCTO video bridge in kilobits per second.
# TYPE jitsi_octo_send_bitrate gauge
jitsi_octo_send_bitrate 0
# HELP jitsi_octo_send_packet_rate The total outgoing packet rate for the OCTO video bridge in packets per second.
# TYPE jitsi_octo_send_packet_rate gauge
jitsi_octo_send_packet_rate 0
# HELP jitsi_outgoing_loss The fraction of lost outgoing RTP packets. This is based on incoming RTCP Receiver Reports, and an attempt to subtract the fraction of packets that were not sent (i.e. were lost before they reached the bridge). Further, this is averaged over all streams of all users as opposed to all packets, so it is not correctly weighted. This is not accurate, but may be a useful metric nonetheless.
# TYPE jitsi_outgoing_loss gauge
jitsi_outgoing_loss 0
# HELP jitsi_p2p_conferences The current number of p2p conferences.
# TYPE jitsi_p2p_conferences gauge
jitsi_p2p_conferences 0
# HELP jitsi_packet_rate_download The total incoming packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_download gauge
jitsi_packet_rate_download 0
# HELP jitsi_packet_rate_upload The total outgoing packet rate for the video bridge in packets per second.
# TYPE jitsi_packet_rate_upload gauge
jitsi_packet_rate_upload 0
# HELP jitsi_participants The current number of participants.
# TYPE jitsi_participants gauge
jitsi_participants 0
# HELP jitsi_receive_only_endpoints The number of endpoints endpoints that have no audio or video, but are not inactive
# TYPE jitsi_receive_only_endpoints gauge
jitsi_receive_only_endpoints 0
# HELP jitsi_rtt_aggregate An average value (in milliseconds) of the RTT across all streams.
# TYPE jitsi_rtt_aggregate gauge
jitsi_rtt_aggregate 0
# HELP jitsi_stress_level Stress Level reported to Jicofo by the videobridge.
# TYPE jitsi_stress_level gauge
jitsi_stress_level 0
# HELP jitsi_threads The number of Java threads that the video bridge is using.
# TYPE jitsi_threads gauge
jitsi_threads 35
# HELP jitsi_total_bytes_received Total bytes received.
# TYPE jitsi_total_bytes_received counter
jitsi_total_bytes_received 0
# HELP jitsi_total_bytes_received_octo The total incoming bit rate for the OCTO video bridge in bytes per second.
# TYPE jitsi_total_bytes_received_octo gauge
jitsi_total_bytes_received_octo 0
# HELP jitsi_total_bytes_sent The number of total bytes sent.
# TYPE jitsi_total_bytes_sent counter
jitsi_total_bytes_sent 0
# HELP jitsi_total_bytes_sent_octo The total outgoing bit rate for the OCTO video bridge in bytes per second.
# TYPE jitsi_total_bytes_sent_octo gauge
jitsi_total_bytes_sent_octo 0
# HELP jitsi_total_colibri_web_socket_messages_received The total number messages received through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_received counter
jitsi_total_colibri_web_socket_messages_received 296
# HELP jitsi_total_colibri_web_socket_messages_sent The total number messages sent through COLIBRI web sockets.
# TYPE jitsi_total_colibri_web_socket_messages_sent counter
jitsi_total_colibri_web_socket_messages_sent 484
# HELP jitsi_total_conference_seconds The sum of the lengths of all completed conferences, in seconds.
# TYPE jitsi_total_conference_seconds counter
jitsi_total_conference_seconds 231
# HELP jitsi_total_conferences_completed The total number of conferences completed on the bridge
# TYPE jitsi_total_conferences_completed counter
jitsi_total_conferences_completed 2
# HELP jitsi_total_conferences_created The total number of conferences created on the bridge.
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created 2
# HELP jitsi_total_data_channel_messages_received The total number messages received through data channels.
# TYPE jitsi_total_data_channel_messages_received counter
jitsi_total_data_channel_messages_received 0
# HELP jitsi_total_data_channel_messages_sent The total number messages sent through data channels.
# TYPE jitsi_total_data_channel_messages_sent counter
jitsi_total_data_channel_messages_sent 0
# HELP jitsi_total_dominant_speaker_changes The cumulative number of times the dominant speaker changed.
# TYPE jitsi_total_dominant_speaker_changes counter
jitsi_total_dominant_speaker_changes 18
# HELP jitsi_total_failed_conferences The total number of failed conferences on the bridge. A conference is marked as failed when all of its channels have failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_failed_conferences counter
jitsi_total_failed_conferences 2
# HELP jitsi_total_ice_failed total number of times ICE failed.
# TYPE jitsi_total_ice_failed counter
jitsi_total_ice_failed 20
# HELP jitsi_total_ice_succeeded Total number of times ICE succeeded.
# TYPE jitsi_total_ice_succeeded counter
jitsi_total_ice_succeeded 0
# HELP jitsi_total_ice_succeeded_relayed Total number of times ICE succeeded.
# TYPE jitsi_total_ice_succeeded_relayed counter
jitsi_total_ice_succeeded_relayed 0
# HELP jitsi_total_ice_succeeded_tcp  Total number of times ICE succeeded over TCP.
# TYPE jitsi_total_ice_succeeded_tcp counter
jitsi_total_ice_succeeded_tcp 0
# HELP jitsi_total_loss_controlled_participant_seconds The total number of participant-seconds that are loss-controlled.
# TYPE jitsi_total_loss_controlled_participant_seconds counter
jitsi_total_loss_controlled_participant_seconds 0
# HELP jitsi_total_loss_degraded_participant_seconds The total number of participant-seconds that are loss-degraded.
# TYPE jitsi_total_loss_degraded_participant_seconds counter
jitsi_total_loss_degraded_participant_seconds 0
# HELP jitsi_total_loss_limited_participant_seconds The total number of participant-seconds that are loss-limited.
# TYPE jitsi_total_loss_limited_participant_seconds counter
jitsi_total_loss_limited_participant_seconds 0
# HELP jitsi_total_packets_dropped_octo The total of dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_dropped_octo gauge
jitsi_total_packets_dropped_octo 0
# HELP jitsi_total_packets_received Total number of packets received
# TYPE jitsi_total_packets_received counter
jitsi_total_packets_received 0
# HELP jitsi_total_packets_received_octo The total of incoming dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_received_octo gauge
jitsi_total_packets_received_octo 0
# HELP jitsi_total_packets_sent The total of sent packets handled by video bridge.
# TYPE jitsi_total_packets_sent gauge
jitsi_total_packets_sent 0
# HELP jitsi_total_packets_sent_octo The total of sent dropped packets handled by the OCTO video bridge.
# TYPE jitsi_total_packets_sent_octo gauge
jitsi_total_packets_sent_octo 0
# HELP jitsi_total_partially_failed_conferences The total number of partially failed conferences on the bridge. A conference is marked as partially failed when some of its channels has failed. A channel is marked as failed if it had no payload activity.
# TYPE jitsi_total_partially_failed_conferences counter
jitsi_total_partially_failed_conferences 0
# HELP jitsi_total_participants The current number of participants.
# TYPE jitsi_total_participants gauge
jitsi_total_participants 20
# HELP jitsi_videochannels The current number of video channels.
# TYPE jitsi_videochannels gauge
jitsi_videochannels 0
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 94.28
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.048576e+06
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 19
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 2.535424e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.61659353594e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.493078016e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 4216
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```