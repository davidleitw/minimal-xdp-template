#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/icmp.h>
#include <linux/in.h>

#define __linux__
#include "../includes/bpf.h"
#include "../includes/bpf_helpers.h"
#include "../includes/bpf_endian.h"

SEC("xdp")
int xdp_root(struct xdp_md *ctx)
{
  void *data = (void *)(long)ctx->data;
  void *data_end = (void *)(long)ctx->data_end;

  // L2
  struct ethhdr *eth = data;
  if (data + sizeof(*eth) > data_end) {
    return XDP_ABORTED;
  }

  // Only Ipv4 supported for this example
  // L3
  if (eth->h_proto != bpf_htons(ETH_P_IP)) {
    return XDP_PASS;
  }
  data += sizeof(*eth);
  
  struct iphdr *ip = data;
  if (data + sizeof(*ip) > data_end) {
    return XDP_ABORTED;
  }

  // Only need ICMP packet
  // L4
  if (ip->protocol != IPPROTO_ICMP) {
    return XDP_PASS;
  }
  data += sizeof(*ip);

  struct icmphdr *icmp = data;
  if (data + sizeof(*icmp) > data_end) {
    return XDP_ABORTED;
  }

  if (bpf_ntohs(icmp->un.echo.sequence) % 2 == 0)
    return XDP_DROP;

  return XDP_PASS;
}

char _license[] SEC("license") = "GPLv2";
