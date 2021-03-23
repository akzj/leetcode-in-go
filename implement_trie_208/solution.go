package implement_trie_208

type Node struct {
	key  string
	next map[int32]*Node
}

type Trie struct {
	nodes map[int32]*Node
}

func (t *Trie) Search(key string) bool {
	var node *Node
	nodes := t.nodes
	for _, ch := range key {
		var ok bool
		if node, ok = nodes[ch]; ok {
			nodes = node.next
			continue
		}
		return false
	}
	return node != nil && node.key == key
}

func (t *Trie) startsWith(key string) bool {
	nodes := t.nodes
	for _, ch := range key {
		if node, ok := nodes[ch]; ok {
			nodes = node.next
			continue
		}
		return false
	}
	return true
}

func (t *Trie) Insert(key string) {
	nodes := t.nodes
	var node *Node
	for _, ch := range key {
		var ok bool
		node, ok = nodes[ch]
		if ok {
			nodes = node.next
		} else {
			node = &Node{
				next: map[int32]*Node{},
			}
			nodes[ch] = node
			nodes = node.next
		}
	}
	if node != nil {
		node.key = key
	}
}

func NewTire() *Trie {
	return &Trie{nodes: map[int32]*Node{}}
}
