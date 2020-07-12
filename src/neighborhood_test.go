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
	if l := len(ancestorsOfAlive); l != 140 {
		t.Errorf("expected 140 ancestors of living, got: %d", l)
	}

	// There are 2^9 = 512 all possible enighborhoods, meaning there are 512-140=372
	// neighborhoods leading to a dead cell.
	if l := len(ancestorsOfDead); l != 372 {
		t.Errorf("expected 372 ancestors of dead, got: %d", l)
	}
}
