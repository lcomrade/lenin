// Copyright 2021 Leonid Maslakov

// License: GPL-3.0-or-later

// This file is part of Lenpaste.

// Lenpaste is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Lenpaste is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Lenpaste.  If not, see <https://www.gnu.org/licenses/>.

package lenin

type remoteError struct {
	Code  int
	Error string
}

type NewReq struct {
	Title      string
	Text       string
	Expiration string
}

type NewResp struct {
	Name string
}

type GetResp struct {
	Name string
	Text string
	Info GetRespInfo
}

type GetRespInfo struct {
	CreateTime int64
	DeleteTime int64
	Title      string
}

type AboutResp struct {
	Exist bool
	Text  string
}

type RulesResp struct {
	Exist bool
	Text  string
}

type VersionResp struct {
	Version   string
	GitTag    string
	GitCommit string
	BuildDate string
}
