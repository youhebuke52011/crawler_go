package engine

import "testing"

func TestConCurrentEngine_Run(t *testing.T) {
	type fields struct {
		Scheduler   Scheduler
		ChanCount   int
		WorkerReady ReadyNotifier
	}
	type args struct {
		sends []Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConCurrentEngine{
				Scheduler:   tt.fields.Scheduler,
				ChanCount:   tt.fields.ChanCount,
				WorkerReady: tt.fields.WorkerReady,
			}
		})
	}
}