package file_content

import "testing"

func Test_CountElements(t *testing.T) {
	type args struct {
		element string
	}
	tests := []struct {
		name    string
		e       *Content
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "bytes",
			e:       New([]byte("hello")),
			args:    args{element: ElementBytes},
			want:    5,
			wantErr: false,
		},
		{
			name:    "runes",
			e:       New([]byte("hello")),
			args:    args{element: ElementRunes},
			want:    5,
			wantErr: false,
		},
		{
			name:    "words",
			e:       New([]byte("hello world")),
			args:    args{element: ElementWords},
			want:    2,
			wantErr: false,
		},
		{
			name:    "lines",
			e:       New([]byte("hello\nworld")),
			args:    args{element: ElementLines},
			want:    1,
			wantErr: false,
		},
		{
			name:    "invalid element",
			e:       New([]byte("hello\nworld")),
			args:    args{element: "invalid"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty content",
			e:       New([]byte("")),
			args:    args{element: ElementLines},
			want:    0,
			wantErr: false,
		},
		{
			name:    "empty content",
			e:       New([]byte("")),
			args:    args{element: ElementWords},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.e
			got, err := e.CountElements(tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("%s: CountElements() got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
