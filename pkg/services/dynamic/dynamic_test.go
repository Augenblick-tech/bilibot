package dynamic

import (
	"testing"

	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Test_checkNew(t *testing.T) {
	type args struct {
		new []bilibot.Dynamic
		old []model.Dynamic
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			name: "test new dynamic (time overlay)",
			args: args{
				new: []bilibot.Dynamic{
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 3,
							},
						},
					},
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 2,
							},
						},
					},
				},
				old: []model.Dynamic{
					{
						PubTS: 2,
					},
					{
						PubTS: 1,
					},
				},
			},
			want:  1,
			want1: 0,
		},
		{
			name: "test new dynamic (no overlay)",
			args: args{
				new: []bilibot.Dynamic{
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 4,
							},
						},
					},
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 3,
							},
						},
					},
				},
				old: []model.Dynamic{
					{
						PubTS: 2,
					},
					{
						PubTS: 1,
					},
				},
			},
			want:  2,
			want1: 0,
		},
		{
			name: "test same dynamic",
			args: args{
				new: []bilibot.Dynamic{
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 3,
							},
						},
					},
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 2,
							},
						},
					},
				},
				old: []model.Dynamic{
					{
						PubTS: 3,
					},
					{
						PubTS: 2,
					},
				},
			},
			want:  0,
			want1: 0,
		},
		{
			name: "test old dynamic (time overlay)",
			args: args{
				new: []bilibot.Dynamic{
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 2,
							},
						},
					},
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 1,
							},
						},
					},
				},
				old: []model.Dynamic{
					{
						PubTS: 3,
					},
					{
						PubTS: 2,
					},
				},
			},
			want:  0,
			want1: 1,
		},
		{
			name: "test old dynamic (no overlay)",
			args: args{
				new: []bilibot.Dynamic{
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 3,
							},
						},
					},
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 2,
							},
						},
					},
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 1,
							},
						},
					},
				},
				old: []model.Dynamic{
					{
						PubTS: 5,
					},
					{
						PubTS: 4,
					},
					{
						PubTS: 3,
					},
				},
			},
			want:  0,
			want1: 2,
		},
		{
			name: "test dynamic (not equal length)",
			args: args{
				new: []bilibot.Dynamic{
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 2,
							},
						},
					},
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 1,
							},
						},
					},
				},
				old: []model.Dynamic{
					{
						PubTS: 3,
					},
				},
			},
			want:  0,
			want1: 1,
		},
		{
			name: "test deleted dynamic (database)",
			args: args{
				new: []bilibot.Dynamic{
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 5,
							},
						},
					},
					{
						Modules: struct {
							Author  bilibot.Author  "json:\"module_author\""
							Content bilibot.Content "json:\"module_dynamic\""
						}{
							Author: bilibot.Author{
								PubTS: 3,
							},
						},
					},
				},
				old: []model.Dynamic{
					{
						PubTS: 8,
					},
					{
						PubTS: 5,
					},
				},
			},
			want:  0,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkNew(tt.args.new, tt.args.old)
			if got != tt.want {
				t.Errorf("checkNew() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkNew() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
