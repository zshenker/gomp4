//
// Copyright 2015 Zac Shenker. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mp4

import (
	"log"
	//"errors"
	"github.com/oikomi/gomp4/util"
)

type TfdtAtom struct {
	Offset                int64
	Size                  int64
	IsFullBox             bool
	Version               uint8
	Flag                  uint32
	BaseMediaDecodeTimeV1 int64
	BaseMediaDecodeTime   int32
	AllBytes              []byte
}

/*
 @version, @flags = stream.ui8, stream.ui24

    if @version == 1
      @base_media_decode_time = stream.si64
    else
      @base_media_decode_time = stream.si32
*/

func TfdtRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {

}
