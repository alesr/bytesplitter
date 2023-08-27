package bytesplitter

import "testing"

func TestSplit(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		data    []byte
		size    ConvertibleToByteSize
		wantLen int
		wantErr bool
	}{
		{
			name:    "Empty data",
			data:    []byte{},
			size:    KB(1),
			wantLen: 0,
		},
		{
			name:    "KB size, even split",
			data:    dummyByteSlice(t, KB(1)),
			size:    KB(1),
			wantLen: 1,
		},
		{
			name:    "KB size, uneven split",
			data:    dummyByteSlice(t, KB(1)+1),
			size:    KB(1),
			wantLen: 2,
		},
		{
			name:    "MB size",
			data:    dummyByteSlice(t, MB(1)),
			size:    MB(1),
			wantLen: 1,
		},
		{
			name:    "Invalid size",
			data:    dummyByteSlice(t, KB(1)),
			size:    KB(0),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Split(tc.data, tc.size)
			if (err != nil) != tc.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if len(got) != tc.wantLen {
				t.Errorf("Split() = %v, want length %v", len(got), tc.wantLen)
			}
		})
	}
}

// dummyByteSlice generates a byte slice of the given size containing 'aaa...aaa'.
func dummyByteSlice(t *testing.T, size ConvertibleToByteSize) []byte {
	t.Helper()

	byteSlice := make([]byte, int(size.toBytes()))

	for i := range byteSlice {
		byteSlice[i] = 'a'
	}
	return byteSlice
}
