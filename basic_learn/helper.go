package main

type Event int

func helper(in <-chan Event,
	out chan<- Event) {
	var q []Event
	for in != nil || len(q) > 0 {
		// Decide whether and what to send.
		var sendOut chan<- Event
		var next Event
		if len(q) > 0 {
			sendOut = out
			next = q[0]
		}
		select {
		case e, ok := <-in:
			if !ok {
				in = nil // stop receiving from in
				break
			}
			q = append(q, e)
		case sendOut <- next:
			q = q[1:]
		}
	}
	close(out)
}
