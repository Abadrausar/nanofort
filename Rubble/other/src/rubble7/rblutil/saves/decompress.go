/*
Copyright 2016 by Milo Christiansen

This software is provided 'as-is', without any express or implied warranty. In
no event will the authors be held liable for any damages arising from the use of
this software.

Permission is granted to anyone to use this software for any purpose, including
commercial applications, and to alter it and redistribute it freely, subject to
the following restrictions:

1. The origin of this software must not be misrepresented; you must not claim
that you wrote the original software. If you use this software in a product, an
acknowledgment in the product documentation would be appreciated but is not
required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source distribution.
*/

package saves

import "compress/zlib"
import "encoding/binary"
import "bytes"
import "errors"

import "dctech/axis"

// Utilities to decompress DF save games.

// Frankly I don't feel like maintaining something like this, maybe later...

ErrNoWorldFile = errors.New("rblutil/saves: Provided directory does not contain a world file.")

// openWorldFile opens world.dat or world.sav from the specified directory.
// The returned reader reads from an in memory buffer containing the file contents
// in *decompressed* form.
// Returns: reader, is .sav, err
func openWorldFile(dir axis.DataSource) (io.Reader, bool, error) {
	var in []byte
	var err error
	sav := false
	if axis.Exists(dir, "world.sav") {
		in, err = axis.ReadAll(dir, "world.sav")
		if err != nil {
			return nil, false, err
		}
		sav = true
	} else if axis.Exists(dir, "world.dat") {
		in, err = axis.ReadAll(dir, "world.dat")
		if err != nil {
			return nil, false, err
		}
	} else {
		return nil, false, ErrNoWorldFile
	}
	
	if len(in) > 8 && in[5] == 1 {
		// Save is compressed
		in := bytes.NewBuffer(in)
		out := new(bytes.Buffer)
		
		var v int32
		err = binary.Read(in, binary.LittleEndian, &v)
		if err != nil {
			return nil, false, err
		}
		
		// Ignore the errors, writes to a byte buffer never fail.
		binary.Write(out, binary.LittleEndian, v) // Version
		binary.Write(out, binary.LittleEndian, int32(0)) // Uncompressed
		
		for in.Len() != 0 {
			var l int32
			err = binary.Read(in, binary.LittleEndian, &l)
			if err != nil {
				return nil, false, err
			}
			buff := new(bytes.Buffer)
			_, err = io.CopyN(buff, in, int64(l))
			if err != nil {
				return nil, false, err
			}
			cr := zlib.NewReader(buff)
			_, err = out.ReadFrom(cr)
			cr.Close()
			if err != nil {
				return nil, false, err
			}
		}
		
		return out, sav, nil
	}
	// Save is not compressed
	return bytes.NewBuffer(in), sav, nil
}

var DFSaveVersions = map[int]string{
	1287: "0.31.01",
	1288: "0.31.02",
	1289: "0.31.03",
	1292: "0.31.04",
	1295: "0.31.05",
	1297: "0.31.06",
	1300: "0.31.08",
	1304: "0.31.09",
	1305: "0.31.10",
	1310: "0.31.11",
	1311: "0.31.12",
	1323: "0.31.13",
	1325: "0.31.14",
	1326: "0.31.15",
	1327: "0.31.16",
	1340: "0.31.17",
	1341: "0.31.18",
	1351: "0.31.19",
	1353: "0.31.20",
	1354: "0.31.21",
	1359: "0.31.22",
	1360: "0.31.23",
	1361: "0.31.24",
	1362: "0.31.25",
	
	1372: "0.34.01",
	1374: "0.34.02",
	1376: "0.34.03",
	1377: "0.34.04",
	1378: "0.34.05",
	1382: "0.34.06",
	1383: "0.34.07",
	1400: "0.34.08",
	1402: "0.34.09",
	1403: "0.34.10",
	1404: "0.34.11",
	
	1441: "0.40.01",
	1442: "0.40.02",
	1443: "0.40.03",
	1444: "0.40.04",
	1445: "0.40.05",
	1446: "0.40.06",
	1448: "0.40.07",
	1449: "0.40.08",
	1451: "0.40.09",
	1452: "0.40.10",
	1456: "0.40.11",
	1459: "0.40.12",
	1462: "0.40.13",
	1469: "0.40.14",
	1470: "0.40.15",
	1471: "0.40.16",
	1472: "0.40.17",
	1473: "0.40.18",
	1474: "0.40.19",
	1477: "0.40.20",
	1478: "0.40.21",
	1479: "0.40.22",
	1480: "0.40.23",
	1481: "0.40.24",
	
	1531: "0.42.01",
	1532: "0.42.02",
	1533: "0.42.03",
	1534: "0.42.04",
	1537: "0.42.05",
	1542: "0.42.06",
}

