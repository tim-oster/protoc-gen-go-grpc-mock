package examplepb

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"io"
	"testing"
)

func TestUnaryCalls(t *testing.T) {
	m := &MockService1Client{}

	m.On_Unary(&Message{Id: 1}).Return(&Message{Id: 2}, nil)
	m.On_Unary(&Message{Id: 2}).Return(&Message{Id: 3}, nil)

	resp, err := m.Unary(context.Background(), &Message{Id: 1})
	require.NoError(t, err)
	assert.Equal(t, &Message{Id: 2}, resp)

	resp, err = m.Unary(context.Background(), &Message{Id: 2})
	require.NoError(t, err)
	assert.Equal(t, &Message{Id: 3}, resp)

	m.Assert_Unary_Called(t, &Message{Id: 1})
	m.Assert_Unary_Called(t, &Message{Id: 2})
	m.Assert_Unary_NumberOfCalls(t, 2)
	m.Assert_Unary_NotCalled(t, &Message{Id: 3})
}

func TestServerStreams(t *testing.T) {
	t.Run("normal stream", func(t *testing.T) {
		m := &MockService1Client{}

		m.On_ServerStream(&Message{Id: 1}).Stream(&Message{Id: 2}, &Message{Id: 3})

		stream, err := m.ServerStream(context.Background(), &Message{Id: 1})
		require.NoError(t, err)

		var msgs []*Message
		for {
			msg, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			require.NoError(t, err)
			msgs = append(msgs, msg)
		}
		assert.Equal(t, []*Message{{Id: 2}, {Id: 3}}, msgs)

		m.Assert_ServerStream_Called(t, &Message{Id: 1})
		m.Assert_ServerStream_NumberOfCalls(t, 1)
		m.Assert_ServerStream_NotCalled(t, &Message{Id: 2})
	})

	t.Run("with header and trailer", func(t *testing.T) {
		m := &MockService1Client{}

		m.On_ServerStream(&Message{Id: 1}).Return(&MockService1_ServerStreamClient{
			RecvHeader:  metadata.Pairs("header", "value"),
			RecvTrailer: metadata.Pairs("trailer", "value"),
			RecvObjects: []*Message{{Id: 2}, {Id: 3}},
		}, nil)

		stream, err := m.ServerStream(context.Background(), &Message{Id: 1})
		require.NoError(t, err)

		header, err := stream.Header()
		require.NoError(t, err)
		assert.Equal(t, metadata.Pairs("header", "value"), header)

		var msgs []*Message
		for {
			msg, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				break
			}
			require.NoError(t, err)
			msgs = append(msgs, msg)
		}
		assert.Equal(t, []*Message{{Id: 2}, {Id: 3}}, msgs)

		assert.Equal(t, metadata.Pairs("trailer", "value"), stream.Trailer())
	})
}
