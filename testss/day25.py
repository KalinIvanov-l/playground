import functools
import networkx as nx
import operator
import re

data = [
    re.split(r": | ", line)
    for line in open("input.txt").read().rstrip().split("\n")
]

G = nx.Graph()
for node, *nodes in data:
    for node2 in nodes:
        G.add_edge(node, node2)

# Uncomment the following two lines to see what edges are to be removed
# nx.draw(G, with_labels=True)
# plt.show()

# Seen surprisingly easily from graph
G.remove_edge("zcj", "rtt")
G.remove_edge("txl", "hxq")
G.remove_edge("gxv", "tpn")

print(
    "Part 1:",
    functools.reduce(operator.mul, (len(g) for g in nx.connected_components(G)), 1),
)
