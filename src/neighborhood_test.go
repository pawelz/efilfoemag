// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package neighborhood

import (
	"fmt"
	"testing"

	"github.com/pawelz/efilfoemag/src/state"
)

func TestParse(t *testing.T) {
	for _, td := range []struct {
		input    string
		expected Neighborhood
		failure  bool
	}{
		{input: "+++++++++", expected: 0x000},
		{input: "++#++##++", expected: 0x04c},
		{input: "++#\n++#\n#++", expected: 0x04c},
		{input: "#########", expected: 0x1ff},
		{input: "####", failure: true},
		{input: "#+++++++++###", failure: true},
	} {
		t.Run(fmt.Sprintf("input %q", td.input), func(t *testing.T) {
			actual, err := Parse(td.input)

			if td.failure {
				if err == nil {
					t.Errorf("expected a failure, got %x", actual)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if actual != td.expected {
				t.Errorf("expected %x, got %x", td.expected, actual)
			}
		})
	}
}

func TestToStr(t *testing.T) {
	for _, td := range []struct {
		input    string
		expected string
		failure  bool
	}{
		{input: "+++++++++", expected: "+++,+++,+++"},
		{input: "++++++++#", expected: "+++,+++,++#"},
		{input: "+++++++#+", expected: "+++,+++,+#+"},
		{input: "+++++++##", expected: "+++,+++,+##"},
		{input: "++++++#++", expected: "+++,+++,#++"},
		{input: "++++++#+#", expected: "+++,+++,#+#"},
		{input: "++++++##+", expected: "+++,+++,##+"},
		{input: "++++++###", expected: "+++,+++,###"},
		{input: "+++++#+++", expected: "+++,++#,+++"},
		{input: "+++++#++#", expected: "+++,++#,++#"},
		{input: "+++++#+#+", expected: "+++,++#,+#+"},
		{input: "+++++#+##", expected: "+++,++#,+##"},
		{input: "+++++##++", expected: "+++,++#,#++"},
		{input: "+++++##+#", expected: "+++,++#,#+#"},
		{input: "+++++###+", expected: "+++,++#,##+"},
		{input: "+++++####", expected: "+++,++#,###"},
		{input: "++++#++++", expected: "+++,+#+,+++"},
		{input: "++++#+++#", expected: "+++,+#+,++#"},
		{input: "++++#++#+", expected: "+++,+#+,+#+"},
		{input: "++++#++##", expected: "+++,+#+,+##"},
		{input: "++++#+#++", expected: "+++,+#+,#++"},
		{input: "++++#+#+#", expected: "+++,+#+,#+#"},
		{input: "++++#+##+", expected: "+++,+#+,##+"},
		{input: "++++#+###", expected: "+++,+#+,###"},
		{input: "++++##+++", expected: "+++,+##,+++"},
		{input: "++++##++#", expected: "+++,+##,++#"},
		{input: "++++##+#+", expected: "+++,+##,+#+"},
		{input: "++++##+##", expected: "+++,+##,+##"},
		{input: "++++###++", expected: "+++,+##,#++"},
		{input: "++++###+#", expected: "+++,+##,#+#"},
		{input: "++++####+", expected: "+++,+##,##+"},
		{input: "++++#####", expected: "+++,+##,###"},
		{input: "+++#+++++", expected: "+++,#++,+++"},
		{input: "+++#++++#", expected: "+++,#++,++#"},
		{input: "+++#+++#+", expected: "+++,#++,+#+"},
		{input: "+++#+++##", expected: "+++,#++,+##"},
		{input: "+++#++#++", expected: "+++,#++,#++"},
		{input: "+++#++#+#", expected: "+++,#++,#+#"},
		{input: "+++#++##+", expected: "+++,#++,##+"},
		{input: "+++#++###", expected: "+++,#++,###"},
		{input: "+++#+#+++", expected: "+++,#+#,+++"},
		{input: "+++#+#++#", expected: "+++,#+#,++#"},
		{input: "+++#+#+#+", expected: "+++,#+#,+#+"},
		{input: "+++#+#+##", expected: "+++,#+#,+##"},
		{input: "+++#+##++", expected: "+++,#+#,#++"},
		{input: "+++#+##+#", expected: "+++,#+#,#+#"},
		{input: "+++#+###+", expected: "+++,#+#,##+"},
		{input: "+++#+####", expected: "+++,#+#,###"},
		{input: "+++##++++", expected: "+++,##+,+++"},
		{input: "+++##+++#", expected: "+++,##+,++#"},
		{input: "+++##++#+", expected: "+++,##+,+#+"},
		{input: "+++##++##", expected: "+++,##+,+##"},
		{input: "+++##+#++", expected: "+++,##+,#++"},
		{input: "+++##+#+#", expected: "+++,##+,#+#"},
		{input: "+++##+##+", expected: "+++,##+,##+"},
		{input: "+++##+###", expected: "+++,##+,###"},
		{input: "+++###+++", expected: "+++,###,+++"},
		{input: "+++###++#", expected: "+++,###,++#"},
		{input: "+++###+#+", expected: "+++,###,+#+"},
		{input: "+++###+##", expected: "+++,###,+##"},
		{input: "+++####++", expected: "+++,###,#++"},
		{input: "+++####+#", expected: "+++,###,#+#"},
		{input: "+++#####+", expected: "+++,###,##+"},
		{input: "+++######", expected: "+++,###,###"},
		{input: "++#++++++", expected: "++#,+++,+++"},
		{input: "++#+++++#", expected: "++#,+++,++#"},
		{input: "++#++++#+", expected: "++#,+++,+#+"},
		{input: "++#++++##", expected: "++#,+++,+##"},
		{input: "++#+++#++", expected: "++#,+++,#++"},
		{input: "++#+++#+#", expected: "++#,+++,#+#"},
		{input: "++#+++##+", expected: "++#,+++,##+"},
		{input: "++#+++###", expected: "++#,+++,###"},
		{input: "++#++#+++", expected: "++#,++#,+++"},
		{input: "++#++#++#", expected: "++#,++#,++#"},
		{input: "++#++#+#+", expected: "++#,++#,+#+"},
		{input: "++#++#+##", expected: "++#,++#,+##"},
		{input: "++#++##++", expected: "++#,++#,#++"},
		{input: "++#++##+#", expected: "++#,++#,#+#"},
		{input: "++#++###+", expected: "++#,++#,##+"},
		{input: "++#++####", expected: "++#,++#,###"},
		{input: "++#+#++++", expected: "++#,+#+,+++"},
		{input: "++#+#+++#", expected: "++#,+#+,++#"},
		{input: "++#+#++#+", expected: "++#,+#+,+#+"},
		{input: "++#+#++##", expected: "++#,+#+,+##"},
		{input: "++#+#+#++", expected: "++#,+#+,#++"},
		{input: "++#+#+#+#", expected: "++#,+#+,#+#"},
		{input: "++#+#+##+", expected: "++#,+#+,##+"},
		{input: "++#+#+###", expected: "++#,+#+,###"},
		{input: "++#+##+++", expected: "++#,+##,+++"},
		{input: "++#+##++#", expected: "++#,+##,++#"},
		{input: "++#+##+#+", expected: "++#,+##,+#+"},
		{input: "++#+##+##", expected: "++#,+##,+##"},
		{input: "++#+###++", expected: "++#,+##,#++"},
		{input: "++#+###+#", expected: "++#,+##,#+#"},
		{input: "++#+####+", expected: "++#,+##,##+"},
		{input: "++#+#####", expected: "++#,+##,###"},
		{input: "++##+++++", expected: "++#,#++,+++"},
		{input: "++##++++#", expected: "++#,#++,++#"},
		{input: "++##+++#+", expected: "++#,#++,+#+"},
		{input: "++##+++##", expected: "++#,#++,+##"},
		{input: "++##++#++", expected: "++#,#++,#++"},
		{input: "++##++#+#", expected: "++#,#++,#+#"},
		{input: "++##++##+", expected: "++#,#++,##+"},
		{input: "++##++###", expected: "++#,#++,###"},
		{input: "++##+#+++", expected: "++#,#+#,+++"},
		{input: "++##+#++#", expected: "++#,#+#,++#"},
		{input: "++##+#+#+", expected: "++#,#+#,+#+"},
		{input: "++##+#+##", expected: "++#,#+#,+##"},
		{input: "++##+##++", expected: "++#,#+#,#++"},
		{input: "++##+##+#", expected: "++#,#+#,#+#"},
		{input: "++##+###+", expected: "++#,#+#,##+"},
		{input: "++##+####", expected: "++#,#+#,###"},
		{input: "++###++++", expected: "++#,##+,+++"},
		{input: "++###+++#", expected: "++#,##+,++#"},
		{input: "++###++#+", expected: "++#,##+,+#+"},
		{input: "++###++##", expected: "++#,##+,+##"},
		{input: "++###+#++", expected: "++#,##+,#++"},
		{input: "++###+#+#", expected: "++#,##+,#+#"},
		{input: "++###+##+", expected: "++#,##+,##+"},
		{input: "++###+###", expected: "++#,##+,###"},
		{input: "++####+++", expected: "++#,###,+++"},
		{input: "++####++#", expected: "++#,###,++#"},
		{input: "++####+#+", expected: "++#,###,+#+"},
		{input: "++####+##", expected: "++#,###,+##"},
		{input: "++#####++", expected: "++#,###,#++"},
		{input: "++#####+#", expected: "++#,###,#+#"},
		{input: "++######+", expected: "++#,###,##+"},
		{input: "++#######", expected: "++#,###,###"},
		{input: "+#+++++++", expected: "+#+,+++,+++"},
		{input: "+#++++++#", expected: "+#+,+++,++#"},
		{input: "+#+++++#+", expected: "+#+,+++,+#+"},
		{input: "+#+++++##", expected: "+#+,+++,+##"},
		{input: "+#++++#++", expected: "+#+,+++,#++"},
		{input: "+#++++#+#", expected: "+#+,+++,#+#"},
		{input: "+#++++##+", expected: "+#+,+++,##+"},
		{input: "+#++++###", expected: "+#+,+++,###"},
		{input: "+#+++#+++", expected: "+#+,++#,+++"},
		{input: "+#+++#++#", expected: "+#+,++#,++#"},
		{input: "+#+++#+#+", expected: "+#+,++#,+#+"},
		{input: "+#+++#+##", expected: "+#+,++#,+##"},
		{input: "+#+++##++", expected: "+#+,++#,#++"},
		{input: "+#+++##+#", expected: "+#+,++#,#+#"},
		{input: "+#+++###+", expected: "+#+,++#,##+"},
		{input: "+#+++####", expected: "+#+,++#,###"},
		{input: "+#++#++++", expected: "+#+,+#+,+++"},
		{input: "+#++#+++#", expected: "+#+,+#+,++#"},
		{input: "+#++#++#+", expected: "+#+,+#+,+#+"},
		{input: "+#++#++##", expected: "+#+,+#+,+##"},
		{input: "+#++#+#++", expected: "+#+,+#+,#++"},
		{input: "+#++#+#+#", expected: "+#+,+#+,#+#"},
		{input: "+#++#+##+", expected: "+#+,+#+,##+"},
		{input: "+#++#+###", expected: "+#+,+#+,###"},
		{input: "+#++##+++", expected: "+#+,+##,+++"},
		{input: "+#++##++#", expected: "+#+,+##,++#"},
		{input: "+#++##+#+", expected: "+#+,+##,+#+"},
		{input: "+#++##+##", expected: "+#+,+##,+##"},
		{input: "+#++###++", expected: "+#+,+##,#++"},
		{input: "+#++###+#", expected: "+#+,+##,#+#"},
		{input: "+#++####+", expected: "+#+,+##,##+"},
		{input: "+#++#####", expected: "+#+,+##,###"},
		{input: "+#+#+++++", expected: "+#+,#++,+++"},
		{input: "+#+#++++#", expected: "+#+,#++,++#"},
		{input: "+#+#+++#+", expected: "+#+,#++,+#+"},
		{input: "+#+#+++##", expected: "+#+,#++,+##"},
		{input: "+#+#++#++", expected: "+#+,#++,#++"},
		{input: "+#+#++#+#", expected: "+#+,#++,#+#"},
		{input: "+#+#++##+", expected: "+#+,#++,##+"},
		{input: "+#+#++###", expected: "+#+,#++,###"},
		{input: "+#+#+#+++", expected: "+#+,#+#,+++"},
		{input: "+#+#+#++#", expected: "+#+,#+#,++#"},
		{input: "+#+#+#+#+", expected: "+#+,#+#,+#+"},
		{input: "+#+#+#+##", expected: "+#+,#+#,+##"},
		{input: "+#+#+##++", expected: "+#+,#+#,#++"},
		{input: "+#+#+##+#", expected: "+#+,#+#,#+#"},
		{input: "+#+#+###+", expected: "+#+,#+#,##+"},
		{input: "+#+#+####", expected: "+#+,#+#,###"},
		{input: "+#+##++++", expected: "+#+,##+,+++"},
		{input: "+#+##+++#", expected: "+#+,##+,++#"},
		{input: "+#+##++#+", expected: "+#+,##+,+#+"},
		{input: "+#+##++##", expected: "+#+,##+,+##"},
		{input: "+#+##+#++", expected: "+#+,##+,#++"},
		{input: "+#+##+#+#", expected: "+#+,##+,#+#"},
		{input: "+#+##+##+", expected: "+#+,##+,##+"},
		{input: "+#+##+###", expected: "+#+,##+,###"},
		{input: "+#+###+++", expected: "+#+,###,+++"},
		{input: "+#+###++#", expected: "+#+,###,++#"},
		{input: "+#+###+#+", expected: "+#+,###,+#+"},
		{input: "+#+###+##", expected: "+#+,###,+##"},
		{input: "+#+####++", expected: "+#+,###,#++"},
		{input: "+#+####+#", expected: "+#+,###,#+#"},
		{input: "+#+#####+", expected: "+#+,###,##+"},
		{input: "+#+######", expected: "+#+,###,###"},
		{input: "+##++++++", expected: "+##,+++,+++"},
		{input: "+##+++++#", expected: "+##,+++,++#"},
		{input: "+##++++#+", expected: "+##,+++,+#+"},
		{input: "+##++++##", expected: "+##,+++,+##"},
		{input: "+##+++#++", expected: "+##,+++,#++"},
		{input: "+##+++#+#", expected: "+##,+++,#+#"},
		{input: "+##+++##+", expected: "+##,+++,##+"},
		{input: "+##+++###", expected: "+##,+++,###"},
		{input: "+##++#+++", expected: "+##,++#,+++"},
		{input: "+##++#++#", expected: "+##,++#,++#"},
		{input: "+##++#+#+", expected: "+##,++#,+#+"},
		{input: "+##++#+##", expected: "+##,++#,+##"},
		{input: "+##++##++", expected: "+##,++#,#++"},
		{input: "+##++##+#", expected: "+##,++#,#+#"},
		{input: "+##++###+", expected: "+##,++#,##+"},
		{input: "+##++####", expected: "+##,++#,###"},
		{input: "+##+#++++", expected: "+##,+#+,+++"},
		{input: "+##+#+++#", expected: "+##,+#+,++#"},
		{input: "+##+#++#+", expected: "+##,+#+,+#+"},
		{input: "+##+#++##", expected: "+##,+#+,+##"},
		{input: "+##+#+#++", expected: "+##,+#+,#++"},
		{input: "+##+#+#+#", expected: "+##,+#+,#+#"},
		{input: "+##+#+##+", expected: "+##,+#+,##+"},
		{input: "+##+#+###", expected: "+##,+#+,###"},
		{input: "+##+##+++", expected: "+##,+##,+++"},
		{input: "+##+##++#", expected: "+##,+##,++#"},
		{input: "+##+##+#+", expected: "+##,+##,+#+"},
		{input: "+##+##+##", expected: "+##,+##,+##"},
		{input: "+##+###++", expected: "+##,+##,#++"},
		{input: "+##+###+#", expected: "+##,+##,#+#"},
		{input: "+##+####+", expected: "+##,+##,##+"},
		{input: "+##+#####", expected: "+##,+##,###"},
		{input: "+###+++++", expected: "+##,#++,+++"},
		{input: "+###++++#", expected: "+##,#++,++#"},
		{input: "+###+++#+", expected: "+##,#++,+#+"},
		{input: "+###+++##", expected: "+##,#++,+##"},
		{input: "+###++#++", expected: "+##,#++,#++"},
		{input: "+###++#+#", expected: "+##,#++,#+#"},
		{input: "+###++##+", expected: "+##,#++,##+"},
		{input: "+###++###", expected: "+##,#++,###"},
		{input: "+###+#+++", expected: "+##,#+#,+++"},
		{input: "+###+#++#", expected: "+##,#+#,++#"},
		{input: "+###+#+#+", expected: "+##,#+#,+#+"},
		{input: "+###+#+##", expected: "+##,#+#,+##"},
		{input: "+###+##++", expected: "+##,#+#,#++"},
		{input: "+###+##+#", expected: "+##,#+#,#+#"},
		{input: "+###+###+", expected: "+##,#+#,##+"},
		{input: "+###+####", expected: "+##,#+#,###"},
		{input: "+####++++", expected: "+##,##+,+++"},
		{input: "+####+++#", expected: "+##,##+,++#"},
		{input: "+####++#+", expected: "+##,##+,+#+"},
		{input: "+####++##", expected: "+##,##+,+##"},
		{input: "+####+#++", expected: "+##,##+,#++"},
		{input: "+####+#+#", expected: "+##,##+,#+#"},
		{input: "+####+##+", expected: "+##,##+,##+"},
		{input: "+####+###", expected: "+##,##+,###"},
		{input: "+#####+++", expected: "+##,###,+++"},
		{input: "+#####++#", expected: "+##,###,++#"},
		{input: "+#####+#+", expected: "+##,###,+#+"},
		{input: "+#####+##", expected: "+##,###,+##"},
		{input: "+######++", expected: "+##,###,#++"},
		{input: "+######+#", expected: "+##,###,#+#"},
		{input: "+#######+", expected: "+##,###,##+"},
		{input: "+########", expected: "+##,###,###"},
		{input: "#++++++++", expected: "#++,+++,+++"},
		{input: "#+++++++#", expected: "#++,+++,++#"},
		{input: "#++++++#+", expected: "#++,+++,+#+"},
		{input: "#++++++##", expected: "#++,+++,+##"},
		{input: "#+++++#++", expected: "#++,+++,#++"},
		{input: "#+++++#+#", expected: "#++,+++,#+#"},
		{input: "#+++++##+", expected: "#++,+++,##+"},
		{input: "#+++++###", expected: "#++,+++,###"},
		{input: "#++++#+++", expected: "#++,++#,+++"},
		{input: "#++++#++#", expected: "#++,++#,++#"},
		{input: "#++++#+#+", expected: "#++,++#,+#+"},
		{input: "#++++#+##", expected: "#++,++#,+##"},
		{input: "#++++##++", expected: "#++,++#,#++"},
		{input: "#++++##+#", expected: "#++,++#,#+#"},
		{input: "#++++###+", expected: "#++,++#,##+"},
		{input: "#++++####", expected: "#++,++#,###"},
		{input: "#+++#++++", expected: "#++,+#+,+++"},
		{input: "#+++#+++#", expected: "#++,+#+,++#"},
		{input: "#+++#++#+", expected: "#++,+#+,+#+"},
		{input: "#+++#++##", expected: "#++,+#+,+##"},
		{input: "#+++#+#++", expected: "#++,+#+,#++"},
		{input: "#+++#+#+#", expected: "#++,+#+,#+#"},
		{input: "#+++#+##+", expected: "#++,+#+,##+"},
		{input: "#+++#+###", expected: "#++,+#+,###"},
		{input: "#+++##+++", expected: "#++,+##,+++"},
		{input: "#+++##++#", expected: "#++,+##,++#"},
		{input: "#+++##+#+", expected: "#++,+##,+#+"},
		{input: "#+++##+##", expected: "#++,+##,+##"},
		{input: "#+++###++", expected: "#++,+##,#++"},
		{input: "#+++###+#", expected: "#++,+##,#+#"},
		{input: "#+++####+", expected: "#++,+##,##+"},
		{input: "#+++#####", expected: "#++,+##,###"},
		{input: "#++#+++++", expected: "#++,#++,+++"},
		{input: "#++#++++#", expected: "#++,#++,++#"},
		{input: "#++#+++#+", expected: "#++,#++,+#+"},
		{input: "#++#+++##", expected: "#++,#++,+##"},
		{input: "#++#++#++", expected: "#++,#++,#++"},
		{input: "#++#++#+#", expected: "#++,#++,#+#"},
		{input: "#++#++##+", expected: "#++,#++,##+"},
		{input: "#++#++###", expected: "#++,#++,###"},
		{input: "#++#+#+++", expected: "#++,#+#,+++"},
		{input: "#++#+#++#", expected: "#++,#+#,++#"},
		{input: "#++#+#+#+", expected: "#++,#+#,+#+"},
		{input: "#++#+#+##", expected: "#++,#+#,+##"},
		{input: "#++#+##++", expected: "#++,#+#,#++"},
		{input: "#++#+##+#", expected: "#++,#+#,#+#"},
		{input: "#++#+###+", expected: "#++,#+#,##+"},
		{input: "#++#+####", expected: "#++,#+#,###"},
		{input: "#++##++++", expected: "#++,##+,+++"},
		{input: "#++##+++#", expected: "#++,##+,++#"},
		{input: "#++##++#+", expected: "#++,##+,+#+"},
		{input: "#++##++##", expected: "#++,##+,+##"},
		{input: "#++##+#++", expected: "#++,##+,#++"},
		{input: "#++##+#+#", expected: "#++,##+,#+#"},
		{input: "#++##+##+", expected: "#++,##+,##+"},
		{input: "#++##+###", expected: "#++,##+,###"},
		{input: "#++###+++", expected: "#++,###,+++"},
		{input: "#++###++#", expected: "#++,###,++#"},
		{input: "#++###+#+", expected: "#++,###,+#+"},
		{input: "#++###+##", expected: "#++,###,+##"},
		{input: "#++####++", expected: "#++,###,#++"},
		{input: "#++####+#", expected: "#++,###,#+#"},
		{input: "#++#####+", expected: "#++,###,##+"},
		{input: "#++######", expected: "#++,###,###"},
		{input: "#+#++++++", expected: "#+#,+++,+++"},
		{input: "#+#+++++#", expected: "#+#,+++,++#"},
		{input: "#+#++++#+", expected: "#+#,+++,+#+"},
		{input: "#+#++++##", expected: "#+#,+++,+##"},
		{input: "#+#+++#++", expected: "#+#,+++,#++"},
		{input: "#+#+++#+#", expected: "#+#,+++,#+#"},
		{input: "#+#+++##+", expected: "#+#,+++,##+"},
		{input: "#+#+++###", expected: "#+#,+++,###"},
		{input: "#+#++#+++", expected: "#+#,++#,+++"},
		{input: "#+#++#++#", expected: "#+#,++#,++#"},
		{input: "#+#++#+#+", expected: "#+#,++#,+#+"},
		{input: "#+#++#+##", expected: "#+#,++#,+##"},
		{input: "#+#++##++", expected: "#+#,++#,#++"},
		{input: "#+#++##+#", expected: "#+#,++#,#+#"},
		{input: "#+#++###+", expected: "#+#,++#,##+"},
		{input: "#+#++####", expected: "#+#,++#,###"},
		{input: "#+#+#++++", expected: "#+#,+#+,+++"},
		{input: "#+#+#+++#", expected: "#+#,+#+,++#"},
		{input: "#+#+#++#+", expected: "#+#,+#+,+#+"},
		{input: "#+#+#++##", expected: "#+#,+#+,+##"},
		{input: "#+#+#+#++", expected: "#+#,+#+,#++"},
		{input: "#+#+#+#+#", expected: "#+#,+#+,#+#"},
		{input: "#+#+#+##+", expected: "#+#,+#+,##+"},
		{input: "#+#+#+###", expected: "#+#,+#+,###"},
		{input: "#+#+##+++", expected: "#+#,+##,+++"},
		{input: "#+#+##++#", expected: "#+#,+##,++#"},
		{input: "#+#+##+#+", expected: "#+#,+##,+#+"},
		{input: "#+#+##+##", expected: "#+#,+##,+##"},
		{input: "#+#+###++", expected: "#+#,+##,#++"},
		{input: "#+#+###+#", expected: "#+#,+##,#+#"},
		{input: "#+#+####+", expected: "#+#,+##,##+"},
		{input: "#+#+#####", expected: "#+#,+##,###"},
		{input: "#+##+++++", expected: "#+#,#++,+++"},
		{input: "#+##++++#", expected: "#+#,#++,++#"},
		{input: "#+##+++#+", expected: "#+#,#++,+#+"},
		{input: "#+##+++##", expected: "#+#,#++,+##"},
		{input: "#+##++#++", expected: "#+#,#++,#++"},
		{input: "#+##++#+#", expected: "#+#,#++,#+#"},
		{input: "#+##++##+", expected: "#+#,#++,##+"},
		{input: "#+##++###", expected: "#+#,#++,###"},
		{input: "#+##+#+++", expected: "#+#,#+#,+++"},
		{input: "#+##+#++#", expected: "#+#,#+#,++#"},
		{input: "#+##+#+#+", expected: "#+#,#+#,+#+"},
		{input: "#+##+#+##", expected: "#+#,#+#,+##"},
		{input: "#+##+##++", expected: "#+#,#+#,#++"},
		{input: "#+##+##+#", expected: "#+#,#+#,#+#"},
		{input: "#+##+###+", expected: "#+#,#+#,##+"},
		{input: "#+##+####", expected: "#+#,#+#,###"},
		{input: "#+###++++", expected: "#+#,##+,+++"},
		{input: "#+###+++#", expected: "#+#,##+,++#"},
		{input: "#+###++#+", expected: "#+#,##+,+#+"},
		{input: "#+###++##", expected: "#+#,##+,+##"},
		{input: "#+###+#++", expected: "#+#,##+,#++"},
		{input: "#+###+#+#", expected: "#+#,##+,#+#"},
		{input: "#+###+##+", expected: "#+#,##+,##+"},
		{input: "#+###+###", expected: "#+#,##+,###"},
		{input: "#+####+++", expected: "#+#,###,+++"},
		{input: "#+####++#", expected: "#+#,###,++#"},
		{input: "#+####+#+", expected: "#+#,###,+#+"},
		{input: "#+####+##", expected: "#+#,###,+##"},
		{input: "#+#####++", expected: "#+#,###,#++"},
		{input: "#+#####+#", expected: "#+#,###,#+#"},
		{input: "#+######+", expected: "#+#,###,##+"},
		{input: "#+#######", expected: "#+#,###,###"},
		{input: "##+++++++", expected: "##+,+++,+++"},
		{input: "##++++++#", expected: "##+,+++,++#"},
		{input: "##+++++#+", expected: "##+,+++,+#+"},
		{input: "##+++++##", expected: "##+,+++,+##"},
		{input: "##++++#++", expected: "##+,+++,#++"},
		{input: "##++++#+#", expected: "##+,+++,#+#"},
		{input: "##++++##+", expected: "##+,+++,##+"},
		{input: "##++++###", expected: "##+,+++,###"},
		{input: "##+++#+++", expected: "##+,++#,+++"},
		{input: "##+++#++#", expected: "##+,++#,++#"},
		{input: "##+++#+#+", expected: "##+,++#,+#+"},
		{input: "##+++#+##", expected: "##+,++#,+##"},
		{input: "##+++##++", expected: "##+,++#,#++"},
		{input: "##+++##+#", expected: "##+,++#,#+#"},
		{input: "##+++###+", expected: "##+,++#,##+"},
		{input: "##+++####", expected: "##+,++#,###"},
		{input: "##++#++++", expected: "##+,+#+,+++"},
		{input: "##++#+++#", expected: "##+,+#+,++#"},
		{input: "##++#++#+", expected: "##+,+#+,+#+"},
		{input: "##++#++##", expected: "##+,+#+,+##"},
		{input: "##++#+#++", expected: "##+,+#+,#++"},
		{input: "##++#+#+#", expected: "##+,+#+,#+#"},
		{input: "##++#+##+", expected: "##+,+#+,##+"},
		{input: "##++#+###", expected: "##+,+#+,###"},
		{input: "##++##+++", expected: "##+,+##,+++"},
		{input: "##++##++#", expected: "##+,+##,++#"},
		{input: "##++##+#+", expected: "##+,+##,+#+"},
		{input: "##++##+##", expected: "##+,+##,+##"},
		{input: "##++###++", expected: "##+,+##,#++"},
		{input: "##++###+#", expected: "##+,+##,#+#"},
		{input: "##++####+", expected: "##+,+##,##+"},
		{input: "##++#####", expected: "##+,+##,###"},
		{input: "##+#+++++", expected: "##+,#++,+++"},
		{input: "##+#++++#", expected: "##+,#++,++#"},
		{input: "##+#+++#+", expected: "##+,#++,+#+"},
		{input: "##+#+++##", expected: "##+,#++,+##"},
		{input: "##+#++#++", expected: "##+,#++,#++"},
		{input: "##+#++#+#", expected: "##+,#++,#+#"},
		{input: "##+#++##+", expected: "##+,#++,##+"},
		{input: "##+#++###", expected: "##+,#++,###"},
		{input: "##+#+#+++", expected: "##+,#+#,+++"},
		{input: "##+#+#++#", expected: "##+,#+#,++#"},
		{input: "##+#+#+#+", expected: "##+,#+#,+#+"},
		{input: "##+#+#+##", expected: "##+,#+#,+##"},
		{input: "##+#+##++", expected: "##+,#+#,#++"},
		{input: "##+#+##+#", expected: "##+,#+#,#+#"},
		{input: "##+#+###+", expected: "##+,#+#,##+"},
		{input: "##+#+####", expected: "##+,#+#,###"},
		{input: "##+##++++", expected: "##+,##+,+++"},
		{input: "##+##+++#", expected: "##+,##+,++#"},
		{input: "##+##++#+", expected: "##+,##+,+#+"},
		{input: "##+##++##", expected: "##+,##+,+##"},
		{input: "##+##+#++", expected: "##+,##+,#++"},
		{input: "##+##+#+#", expected: "##+,##+,#+#"},
		{input: "##+##+##+", expected: "##+,##+,##+"},
		{input: "##+##+###", expected: "##+,##+,###"},
		{input: "##+###+++", expected: "##+,###,+++"},
		{input: "##+###++#", expected: "##+,###,++#"},
		{input: "##+###+#+", expected: "##+,###,+#+"},
		{input: "##+###+##", expected: "##+,###,+##"},
		{input: "##+####++", expected: "##+,###,#++"},
		{input: "##+####+#", expected: "##+,###,#+#"},
		{input: "##+#####+", expected: "##+,###,##+"},
		{input: "##+######", expected: "##+,###,###"},
		{input: "###++++++", expected: "###,+++,+++"},
		{input: "###+++++#", expected: "###,+++,++#"},
		{input: "###++++#+", expected: "###,+++,+#+"},
		{input: "###++++##", expected: "###,+++,+##"},
		{input: "###+++#++", expected: "###,+++,#++"},
		{input: "###+++#+#", expected: "###,+++,#+#"},
		{input: "###+++##+", expected: "###,+++,##+"},
		{input: "###+++###", expected: "###,+++,###"},
		{input: "###++#+++", expected: "###,++#,+++"},
		{input: "###++#++#", expected: "###,++#,++#"},
		{input: "###++#+#+", expected: "###,++#,+#+"},
		{input: "###++#+##", expected: "###,++#,+##"},
		{input: "###++##++", expected: "###,++#,#++"},
		{input: "###++##+#", expected: "###,++#,#+#"},
		{input: "###++###+", expected: "###,++#,##+"},
		{input: "###++####", expected: "###,++#,###"},
		{input: "###+#++++", expected: "###,+#+,+++"},
		{input: "###+#+++#", expected: "###,+#+,++#"},
		{input: "###+#++#+", expected: "###,+#+,+#+"},
		{input: "###+#++##", expected: "###,+#+,+##"},
		{input: "###+#+#++", expected: "###,+#+,#++"},
		{input: "###+#+#+#", expected: "###,+#+,#+#"},
		{input: "###+#+##+", expected: "###,+#+,##+"},
		{input: "###+#+###", expected: "###,+#+,###"},
		{input: "###+##+++", expected: "###,+##,+++"},
		{input: "###+##++#", expected: "###,+##,++#"},
		{input: "###+##+#+", expected: "###,+##,+#+"},
		{input: "###+##+##", expected: "###,+##,+##"},
		{input: "###+###++", expected: "###,+##,#++"},
		{input: "###+###+#", expected: "###,+##,#+#"},
		{input: "###+####+", expected: "###,+##,##+"},
		{input: "###+#####", expected: "###,+##,###"},
		{input: "####+++++", expected: "###,#++,+++"},
		{input: "####++++#", expected: "###,#++,++#"},
		{input: "####+++#+", expected: "###,#++,+#+"},
		{input: "####+++##", expected: "###,#++,+##"},
		{input: "####++#++", expected: "###,#++,#++"},
		{input: "####++#+#", expected: "###,#++,#+#"},
		{input: "####++##+", expected: "###,#++,##+"},
		{input: "####++###", expected: "###,#++,###"},
		{input: "####+#+++", expected: "###,#+#,+++"},
		{input: "####+#++#", expected: "###,#+#,++#"},
		{input: "####+#+#+", expected: "###,#+#,+#+"},
		{input: "####+#+##", expected: "###,#+#,+##"},
		{input: "####+##++", expected: "###,#+#,#++"},
		{input: "####+##+#", expected: "###,#+#,#+#"},
		{input: "####+###+", expected: "###,#+#,##+"},
		{input: "####+####", expected: "###,#+#,###"},
		{input: "#####++++", expected: "###,##+,+++"},
		{input: "#####+++#", expected: "###,##+,++#"},
		{input: "#####++#+", expected: "###,##+,+#+"},
		{input: "#####++##", expected: "###,##+,+##"},
		{input: "#####+#++", expected: "###,##+,#++"},
		{input: "#####+#+#", expected: "###,##+,#+#"},
		{input: "#####+##+", expected: "###,##+,##+"},
		{input: "#####+###", expected: "###,##+,###"},
		{input: "######+++", expected: "###,###,+++"},
		{input: "######++#", expected: "###,###,++#"},
		{input: "######+#+", expected: "###,###,+#+"},
		{input: "######+##", expected: "###,###,+##"},
		{input: "#######++", expected: "###,###,#++"},
		{input: "#######+#", expected: "###,###,#+#"},
		{input: "########+", expected: "###,###,##+"},
		{input: "#########", expected: "###,###,###"},
	} {
		t.Run(fmt.Sprintf("input %q", td.input), func(t *testing.T) {
			p, err := Parse(td.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if actual := p.ToStr(); actual != td.expected {
				t.Errorf("expected %q got %q", td.expected, actual)
			}
		})
	}
}

// testName assembles a nice human readable name for the Matches() test.
func testName(t *testing.T, n string, k string, dist int, side Side) string {
	t.Helper()
	sanitize := func(n string) string {
		p, err := Parse(n)
		if err != nil {
			t.Fatalf(err.Error())
		}
		return p.ToStr()
	}
	return fmt.Sprintf("%q matches %q at distance %d side %s", sanitize(n), sanitize(k), dist, side.ToStr())
}

func TestMask(t *testing.T) {
	for _, td := range []struct {
		input    Side
		expected uint16
	}{
		{
			input:    NW,
			expected: 0x1ff - 0x100,
		},
		{
			input:    N,
			expected: 0x1ff - 0x080,
		},
		{
			input:    NE,
			expected: 0x1ff - 0x040,
		},
		{
			input:    W,
			expected: 0x1ff - 0x020,
		},
		{
			input:    C,
			expected: 0x1ff - 0x010,
		},
		{
			input:    E,
			expected: 0x1ff - 0x008,
		},
		{
			input:    SW,
			expected: 0x1ff - 0x004,
		},
		{
			input:    S,
			expected: 0x1ff - 0x002,
		},
		{
			input:    SE,
			expected: 0x1ff - 0x001,
		},
	} {
		t.Run(td.input.ToStr(), func(t *testing.T) {
			if got := mask(td.input); got != td.expected {
				t.Errorf("expected %x got %x", td.expected, got)
			}
		})
	}
}

func TestSet(t *testing.T) {
	for _, td := range []struct {
		n        string
		side     Side
		value    state.State
		expected string
	}{
		{
			n: `+++
			    +++
			    +++`,
			side:  NW,
			value: state.Alive,
			expected: `#++
			           +++
			           +++`,
		},
		{
			n: `+++
			    +++
			    +++`,
			side:  N,
			value: state.Alive,
			expected: `+#+
			           +++
			           +++`,
		},
		{
			n: `+++
			    +++
			    +++`,
			side:  NE,
			value: state.Alive,
			expected: `++#
			           +++
			           +++`,
		},
		{
			n: `+++
			    +++
			    +++`,
			side:  W,
			value: state.Alive,
			expected: `+++
			           #++
			           +++`,
		},
		{
			n: `+++
			    +++
			    +++`,
			side:  C,
			value: state.Alive,
			expected: `+++
			           +#+
			           +++`,
		},
		{
			n: `+++
			    +++
			    +++`,
			side:  E,
			value: state.Alive,
			expected: `+++
			           ++#
			           +++`,
		},
		{
			n: `+++
			    +++
			    +++`,
			side:  SW,
			value: state.Alive,
			expected: `+++
			           +++
			           #++`,
		},
		{
			n: `+++
			    +++
			    +++`,
			side:  S,
			value: state.Alive,
			expected: `+++
			           +++
			           +#+`,
		},
		{
			n: `+++
			    +++
			    +++`,
			side:  SE,
			value: state.Alive,
			expected: `+++
			           +++
			           ++#`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  NW,
			value: state.Dead,
			expected: `+##
			           ###
			           ###`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  N,
			value: state.Dead,
			expected: `#+#
			           ###
			           ###`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  NE,
			value: state.Dead,
			expected: `##+
			           ###
			           ###`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  W,
			value: state.Dead,
			expected: `###
			           +##
			           ###`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  C,
			value: state.Dead,
			expected: `###
			           #+#
			           ###`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  E,
			value: state.Dead,
			expected: `###
			           ##+
			           ###`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  SW,
			value: state.Dead,
			expected: `###
			           ###
			           +##`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  S,
			value: state.Dead,
			expected: `###
			           ###
			           #+#`,
		},
		{
			n: `###
			    ###
			    ###`,
			side:  SE,
			value: state.Dead,
			expected: `###
			           ###
			           ##+`,
		},
	} {
		actual, err := Parse(td.n)
		if err != nil {
			t.Fatalf(err.Error())
		}
		want, err := Parse(td.expected)
		if err != nil {
			t.Fatalf(err.Error())
		}
		t.Run(fmt.Sprintf("%q.Set(%s, %s)", actual.ToStr(), td.side.ToStr(), td.value.ToStr()), func(t *testing.T) {
			(&actual).Set(td.side, td.value)
			if actual != want {
				t.Errorf("expected %q got %q", want.ToStr(), actual.ToStr())
			}
		})
	}
}

func TestMatches(t *testing.T) {
	for _, td := range []struct {
		n            string
		k            string
		matchesDist1 []Side
		matchesDist2 []Side
	}{
		{
			n: `###
			    ###
					###`,
			k: `###
			    ###
					###`,
			matchesDist1: []Side{NW, N, NE, W, E, SW, S, SE},
			matchesDist2: []Side{NW, N, NE, W, E, SW, S, SE},
		},
		{
			n: `###
			    ++#
					#+#`,
			k: `+##
			    +##
					###`,
			matchesDist1: []Side{SE},
			matchesDist2: []Side{NW, N, NE, SW},
		},
		{
			n: `###
			    ##+
					###`,
			k: `###
			    +##
					###`,
			matchesDist1: []Side{NW, W, SW},
			matchesDist2: []Side{NW, N, NE, W, E, SW, S, SE},
		},
		{
			n: `###
			    +##
					###`,
			k: `###
			    ##+
					###`,
			matchesDist1: []Side{NE, E, SE},
			matchesDist2: []Side{NW, N, NE, W, E, SW, S, SE},
		},
		{
			n: `#+#
			    ###
					###`,
			k: `###
			    ###
					#+#`,
			matchesDist1: []Side{SE, S, SW},
			matchesDist2: []Side{NW, N, NE, W, E, SW, S, SE},
		},
		{
			n: `###
			    ###
					#+#`,
			k: `#+#
			    ###
					###`,
			matchesDist1: []Side{NW, N, NE},
			matchesDist2: []Side{NW, N, NE, W, E, SW, S, SE},
		},
	} {
		matchSet := map[int]map[Side]bool{
			1: make(map[Side]bool),
			2: make(map[Side]bool),
		}
		for _, side := range td.matchesDist1 {
			matchSet[1][side] = true
		}
		for _, side := range td.matchesDist2 {
			matchSet[2][side] = true
		}

		for _, dist := range []int{1, 2} {
			for _, side := range []Side{NW, N, NE, W, E, SW, S, SE} {
				t.Run(testName(t, td.n, td.k, dist, side), func(t *testing.T) {
					expected := matchSet[dist][side]
					parsedN, err := Parse(td.n)
					if err != nil {
						t.Fatalf("unable to parse testdata n: %v", err)
					}
					parsedK, err := Parse(td.k)
					if err != nil {
						t.Fatalf("unable to parse testdata k: %v", err)
					}
					actual, err := parsedN.Matches(parsedK, dist, side)
					if err != nil {
						t.Errorf("unexpected error: %v", err)
					}
					if expected != actual {
						t.Errorf("want %v got %v", expected, actual)
					}
				})
			}
		}
	}
}

func TestAncestors(t *testing.T) {
	// There should be 140 neighborhoods leading to alive cell in the next turn.
	// Those neighborhoods are:
	// 1. living cell in the middle, exactly three living cells around. There are
	//    8 cells around, that means there are (8 3) = 8! / (8 - 3)!3! = 56 of those.
	// 2. living cell in the middle, exactly two living cells around. There are
	//    (8 2) = 28 of those.
	// 3. dead cell in the middle, exactly three living cells around, so this is
	//    again 56.
	// So in total there are 56+28+56 = 140 neighborhoods leading to an alive cell.
	t.Run("ancestors of alive", func(t *testing.T) {
		var counter int

		for iter := ancestorsOfAlive.iterate(); iter.hasNext(); _ = iter.getNext() {
			counter++
		}

		if counter != 140 {
			t.Errorf("expected 140 ancestors of living, got: %d", counter)
		}
	})

	// There are 2^9 = 512 all possible enighborhoods, meaning there are 512-140=372
	// neighborhoods leading to a dead cell.
	t.Run("ancestors of alive", func(t *testing.T) {
		var counter int

		for iter := ancestorsOfDead.iterate(); iter.hasNext(); _ = iter.getNext() {
			counter++
		}

		if counter != 372 {
			t.Errorf("expected 372 ancestors of dead, got: %d", counter)
		}
	})
}

func TestSetAddString(t *testing.T) {
	for _, td := range []struct{
		given []string
		want string
	} {
		{
			given: []string{},
			want: "[]",
		},
		{
			given: []string{"### +++ ###"},
			want: "[\"###,+++,###\"]",
		},
		{
			given: []string{
				"##+ #++ +#+",
				"+++ #++ +#+",
				"##+ ### ##+",
				"+#+ #++ +++",
			},
			want: `["+++,#++,+#+", "+#+,#++,+++", "##+,#++,+#+", "##+,###,##+"]`,
		},
		{
			given: []string{
				"##+ #++ +#+",
				"##+ #++ +#+",
				"+#+ #++ +++",
				"+#+ #++ +++",
			},
			want: `["+#+,#++,+++", "##+,#++,+#+"]`,
		},
	} {
		t.Run(td.want, func(t *testing.T) {
			a := &Set{}
			for _, g := range td.given {
				parsed, err := Parse(g)
				if err != nil {
					t.Fatal(err)
				}
				err = a.Add(parsed)
				if err != nil {
					t.Fatal(err)
				}
			}
			if got := a.String(); td.want != got {
				t.Errorf("want %s got %s", td.want, got)
			}
		})
	}
}

func TestSetEquals(t *testing.T) {
	a := &Set{}
	b := &Set{}
	for _, n := range []string{
		"##+ #++ +#+",
		"+++ #++ +#+",
		"##+ ### ##+",
		"+#+ #++ +++",
	} {
		parsed, err := Parse(n)
		if err != nil {
			t.Errorf("unable to parse testdata %q: %v", n, err)
		}
	  if err = a.Add(parsed); err != nil {
		  t.Errorf("unable to add %q to 'a'; want no error, got %v", n, err)
	  }
	  if err = b.Add(parsed); err != nil {
		  t.Errorf("unable to add %q to 'a'; want no error, got %v", n, err)
	  }
	}
	if !Equals(a, b) {
		t.Errorf("expected equal sets, got left=%v, right=%v", a, b)
	}
	n := "### +++ ###"
	parsed, err := Parse(n)
	if err != nil {
		t.Errorf("unable to parse testdata %q: %v", n, err)
	}
	a.Add(parsed)
	if Equals(a, b) {
		t.Errorf("expected not equal sets, got left=%v, right=%v", a, b)
	}
}

type pairOfSets struct {
	left []string
	right []string
}

func TestShiftIntersect(t *testing.T) {
	for _, td := range []struct{
		name string
		given pairOfSets
		side Side
		want pairOfSets
	} {
		{
			name: "empty",
			side: E,
		},
		{
			name: "one matching element",
			given: pairOfSets{
				left: []string{"+#+,++#, #++"},
				right: []string{"#++,+##, +++"},
			},
			side: E,
			want: pairOfSets{
				left: []string{"+#+,++#, #++"},
				right: []string{"#++,+##, +++"},
			},
		},
		{
			name: "no matching elements",
			given: pairOfSets{
				left: []string{"+#+,++#, #++"},
				right: []string{"#++,++#, +++"},
			},
		},
		{
			name: "multiple matching elements",
			given: pairOfSets{
				left: []string{"+#+,++#,#++", "++#,++#,##+", "###,###,###"},
				right: []string{"#++,+##,+++", "+##,+#+,#+#", "+++,+++,+++"},
			},
			side: E,
			want: pairOfSets{
				left: []string{"+#+,++#,#++", "++#,++#,##+"},
				right: []string{"#++,+##,+++", "+##,+#+,#+#"},
			},
		},
	} {
		parseList := func(input []string, t *testing.T) *Set {
			t.Helper()
			rv := &Set{}
			for _, g := range input {
				parsed, err := Parse(g)
				if err != nil {
					t.Fatal(err)
				}
				if err = rv.Add(parsed); err != nil {
					t.Fatal(err)
				}
			}
			return rv
		}
		t.Run(td.name, func(t *testing.T) {
			givenL := parseList(td.given.left, t)
			givenR := parseList(td.given.right, t)
			wantL := parseList(td.want.left, t)
			wantR := parseList(td.want.right, t)
			gotL, gotR, err := ShiftIntersect(givenL, givenR, td.side)
			if err != nil {
				t.Fatalf("ShiftIntersect returned error: %v", err)
			}
			if !Equals(wantL, gotL) {
				t.Errorf("the left set is invalid: want %v, got %v", wantL, gotL)
			}
			if !Equals(wantR, gotR) {
				t.Errorf("the right set is invalid: want %v, got %v", wantR, gotR)
			}
		})
	}
}
