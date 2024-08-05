package functions

import (
	"os"
	"testing"
)

func TestVerifyFileChecksum(t *testing.T) {
	type args struct {
		banner, filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Valid standard.txt file",
			args:    args{banner: "standard.txt", filename: "../standard.txt"},
			wantErr: false,
		},
		{
			name:    "Valid shadow.txt file",
			args:    args{banner: "shadow.txt", filename: "../shadow.txt"},
			wantErr: false,
		},
		{
			name:    "Valid thinkertoy.txt file",
			args:    args{banner: "thinkertoy.txt", filename: "../thinkertoy.txt"},
			wantErr: false,
		},
		{
			name:    "Non-existent file",
			args:    args{filename: "nonexistent.txt"},
			wantErr: true,
		},
		{
			name:    "Corrupted file",
			args:    args{filename: "corrupted.txt"},
			wantErr: true,
		},
	}

	// Create a temporary corrupted file for testing
	corruptedFile, _ := os.Create("corrupted.txt")
	corruptedFile.WriteString("Corrupted content")
	corruptedFile.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyFileChecksum(tt.args.banner, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("VerifyFileChecksum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// Clean up the temporary corrupted file
	os.Remove("corrupted.txt")
}
