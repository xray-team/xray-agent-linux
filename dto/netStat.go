package dto

/*
netstat -s = content of file /proc/net/netstat + content of file /proc/net/snmp

	/proc/net/netstat
		TcpExt:
			SyncookiesFailed			- invalid SYN cookies received
			EmbryonicRsts				- resets received for embryonic SYN_RECV sockets
			PruneCalled 				- packets pruned from receive queue because of socket buffer overrun
			OutOfWindowIcmps			- ICMP packets dropped because they were graphite-of-window
			LockDroppedIcmps			- ICMP packets dropped because socket was locked
			TW 							- TCP sockets finished time wait in fast timer
			PAWSEstab					- packetes rejected in established connections because of timestamp
			DelayedACKs					- delayed acks sent
			DelayedACKLocked			- delayed acks further delayed because of locked socket
			DelayedACKLost				- quick ack mode was activated 'DelayedACKLost' times
			TCPPrequeued				- packets directly queued to recvmsg prequeue
			TCPDirectCopyFromBacklog 	- bytes directly in process context from backlog
			TCPDirectCopyFromPrequeue	- bytes directly received in process context from prequeue
			TCPPrequeueDropped			- packets dropped from prequeue
			TCPHPHits					- packet headers predicted
			TCPHPHitsToUser				- packets header predicted and directly queued to user
			TCPPureAcks					- acknowledgments not containing data payload received
			TCPHPAcks					- predicted acknowledgments
			TCPSackRecovery				- times recovered from packet loss by selective acknowledgements
			TCPSACKReneging				- bad SACK blocks received
			TCPFACKReorder				- detected reordering 'TCPFACKReorder' times using FACK
			TCPSACKReorder				- detected reordering 'TCPSACKReorder' times using SACK
			TCPTSReorder				- detected reordering 'TCPTSReorder' times using time stamp
			TCPFullUndo					- congestion windows fully recovered without slow start
			TCPPartialUndo				- congestion windows partially recovered using Hoe heuristic
			TCPDSACKUndo				- congestion windows recovered without slow start by DSACK
			TCPLossUndo					- congestion windows recovered without slow start after partial ack
			TCPLoss						- TCP data loss events
			TCPSackFailures				- timeouts after SACK recovery
			TCPLossFailures				- timeouts in loss state
			TCPFastRetrans				- fast retransmits
			TCPForwardRetrans			- forward retransmits
			TCPSlowStartRetrans			- retransmits in slow start
			TCPTimeouts					- other TCP timeouts
			TCPSackRecoveryFail			- SACK retransmits failed
			TCPDSACKOldSent				- DSACKs sent for old packets
			TCPDSACKOfoSent				- DSACKs sent for graphite of order packets
			TCPDSACKRecv				- DSACKs received
			TCPDSACKOfoRecv				- DSACKs for graphite of order packets received
			TCPRcvCollapsed				- packets collapsed in receive queue due to low socket buffer
			TCPAbortOnData				- connections reset due to unexpected data
			TCPAbortOnClose				- connections reset due to early user close
			TCPAbortOnTimeout			- connections aborted due to timeout

	/proc/net/snmp
		IP:
			InReceives			- total packets received
			InHdrErrors			- packets received with invalid headers
			InAddrErrors		- packets received with invalid addresses
			InDiscards			- incoming packets discarded
			InDelivers			- incoming packets delivered
			OutRequests			- requests sent graphite
			OutDiscards			- outgoing packets dropped
			OutNoRoutes			- dropped because of missing route
		Icmp:
			InMsgs				- ICMP messages received
			InErrors			- input ICMP message failed
			InDestUnreachs		- received ICMP message type: destination unreachable
			InTimeExcds			- received ICMP message type: timeout in transit
			InEchos				- received ICMP message type: echo requests
			InEchoReps			- received ICMP message type: echo replies
			InTimestampReps		- received ICMP message type: timestamp reply
			OutMsgs				- ICMP messages sent
			OutErrors			- ICMP messages failed
			OutDestUnreachs		- sent graphite ICMP messages type: destination unreachable
			OutEchos			- sent graphite ICMP messages type: echo requests
			OutEchoReps			- sent graphite ICMP messages type: echo replies
		Tcp:
			ActiveOpens			- active connection openings
			PassiveOpens		- passive connection openings
			AttemptFails		- failed connection attempts
			EstabResets			- connection resets received
			CurrEstab			- connections established
			InSegs				- segments received
			OutSegs				- segments sent
			RetransSegs			- segments retransmitted
			InErrs				- bad segments received
			OutRsts				- resets sent
		Udp:
			InDatagrams			- packets received
			NoPorts				- packets to unknown port received
			InErrors			- packet receive errors
			OutDatagrams		- packets sent
			RcvbufErrors		- receive buffer errors
			SndbufErrors		- send buffer errors

*/

// Netstat describes the content of the files /proc/net/netstat, /proc/net/snmp (/proc/$PID/net/netstat, /proc/$PID/net/snmp)
type Netstat struct {
	Ext map[string]map[string]int64
}

// NetSNMP6 describes the content of the file /proc/net/snmp6 (/proc/$PID/net/snmp6)
type NetSNMP6 struct {
	Counters map[string]int64
}
