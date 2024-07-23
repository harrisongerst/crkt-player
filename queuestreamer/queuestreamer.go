package queuestreamer

import (
	"github.com/gopxl/beep"
)

type QueueStreamer struct {
	Streamers []beep.StreamSeekCloser
	Current   int
	Paused    bool
	Loading   bool
	Start     int
	End       int
}

// Err implements beep.Streamer. needed this to be beep.Streamer TODO figure out what i need to do with this
func (q *QueueStreamer) Err() error {
	return nil
}

func (q *QueueStreamer) Add(Streamers ...beep.StreamSeekCloser) {
	q.Streamers = append(q.Streamers, Streamers...)
}

func (q *QueueStreamer) Stream(samples [][2]float64) (n int, ok bool) {
	// We use the filled variable to track how many samples we've
	// successfully filled already. We loop until all samples are filled.
	filled := 0
	for filled < len(samples) {
		// There are no Streamers in the QueueStreamer, so we stream silence.
		if len(q.Streamers) == 0 || q.Paused || q.Current < 0 {
			for i := range samples[filled:] {
				samples[i][0] = 0
				samples[i][1] = 0
			}
			break
		}

		// We stream from the first streamer in the QueueStreamer.
		n, ok := q.Streamers[q.Current].Stream(samples[filled:])
		// If it's drained, we pop it from the QueueStreamer, thus continuing with
		// the next streamer.
		if !ok {
			streamer := q.Streamers[q.Current]
			streamer.Seek(0)
			q.Current++
			if q.Current == len(q.Streamers) {
				q.Current = 0
			}

		}
		// We update the number of filled samples.
		filled += n
	}
	return len(samples), true
}
