package proc_test

import (
	"reflect"
	"testing"

	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
	"xray-agent-linux/proc"
)

func Test_netSNMP6DataSource_GetData(t *testing.T) {
	logger.Init("")

	tests := []struct {
		caseDescription string
		filePath        string
		want            *dto.NetSNMP6
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testfiles/netNetstat/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "snmp6",
			filePath:        "./testfiles/netNetstat/snmp6-kernel5.4.0-40-generic",
			want: &dto.NetSNMP6{
				Counters: map[string]int64{
					"Ip6InReceives":                  1,
					"Ip6InHdrErrors":                 2,
					"Ip6InTooBigErrors":              3,
					"Ip6InNoRoutes":                  4,
					"Ip6InAddrErrors":                5,
					"Ip6InUnknownProtos":             6,
					"Ip6InTruncatedPkts":             7,
					"Ip6InDiscards":                  8,
					"Ip6InDelivers":                  9,
					"Ip6OutForwDatagrams":            10,
					"Ip6OutRequests":                 11,
					"Ip6OutDiscards":                 12,
					"Ip6OutNoRoutes":                 13,
					"Ip6ReasmTimeout":                14,
					"Ip6ReasmReqds":                  15,
					"Ip6ReasmOKs":                    16,
					"Ip6ReasmFails":                  17,
					"Ip6FragOKs":                     18,
					"Ip6FragFails":                   19,
					"Ip6FragCreates":                 20,
					"Ip6InMcastPkts":                 23,
					"Ip6OutMcastPkts":                53,
					"Ip6InOctets":                    3625,
					"Ip6OutOctets":                   5221,
					"Ip6InMcastOctets":               3147,
					"Ip6OutMcastOctets":              5123,
					"Ip6InBcastOctets":               0,
					"Ip6OutBcastOctets":              0,
					"Ip6InNoECTPkts":                 30,
					"Ip6InECT1Pkts":                  0,
					"Ip6InECT0Pkts":                  0,
					"Ip6InCEPkts":                    0,
					"Icmp6InMsgs":                    5,
					"Icmp6InErrors":                  0,
					"Icmp6OutMsgs":                   30,
					"Icmp6OutErrors":                 0,
					"Icmp6InCsumErrors":              0,
					"Icmp6InDestUnreachs":            0,
					"Icmp6InPktTooBigs":              0,
					"Icmp6InTimeExcds":               0,
					"Icmp6InParmProblems":            0,
					"Icmp6InEchos":                   0,
					"Icmp6InEchoReplies":             0,
					"Icmp6InGroupMembQueries":        0,
					"Icmp6InGroupMembResponses":      0,
					"Icmp6InGroupMembReductions":     0,
					"Icmp6InRouterSolicits":          0,
					"Icmp6InRouterAdvertisements":    0,
					"Icmp6InNeighborSolicits":        0,
					"Icmp6InNeighborAdvertisements":  0,
					"Icmp6InRedirects":               0,
					"Icmp6InMLDv2Reports":            5,
					"Icmp6OutDestUnreachs":           0,
					"Icmp6OutPktTooBigs":             0,
					"Icmp6OutTimeExcds":              0,
					"Icmp6OutParmProblems":           0,
					"Icmp6OutEchos":                  0,
					"Icmp6OutEchoReplies":            0,
					"Icmp6OutGroupMembQueries":       0,
					"Icmp6OutGroupMembResponses":     0,
					"Icmp6OutGroupMembReductions":    0,
					"Icmp6OutRouterSolicits":         15,
					"Icmp6OutRouterAdvertisements":   0,
					"Icmp6OutNeighborSolicits":       1,
					"Icmp6OutNeighborAdvertisements": 0,
					"Icmp6OutRedirects":              0,
					"Icmp6OutMLDv2Reports":           14,
					"Icmp6InType143":                 5,
					"Icmp6OutType133":                15,
					"Icmp6OutType135":                1,
					"Icmp6OutType143":                14,
					"Udp6InDatagrams":                24,
					"Udp6NoPorts":                    0,
					"Udp6InErrors":                   0,
					"Udp6OutDatagrams":               25,
					"Udp6RcvbufErrors":               0,
					"Udp6SndbufErrors":               0,
					"Udp6InCsumErrors":               0,
					"Udp6IgnoredMulti":               0,
					"UdpLite6InDatagrams":            0,
					"UdpLite6NoPorts":                0,
					"UdpLite6InErrors":               0,
					"UdpLite6OutDatagrams":           0,
					"UdpLite6RcvbufErrors":           0,
					"UdpLite6SndbufErrors":           0,
					"UdpLite6InCsumErrors":           0,
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			netSNMP6DataSource := proc.NewNetSNMP6DataSource(tt.filePath, "")
			got, err := netSNMP6DataSource.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNetSnmp6() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNetSnmp6() got = %v, want %v", got, tt.want)
			}
		})
	}
}
