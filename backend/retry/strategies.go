// Copyright 2024 BiocryptoID
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

package retry

import (
	"math"
	"math/rand"
	"time"
)

type Strategy interface {
	Duration(attempt int) time.Duration
}

type ExponentialStrategy struct {
	Min       time.Duration
	Max       time.Duration
	MaxJitter time.Duration
}

func (e *ExponentialStrategy) Duration(attempt int) time.Duration {
	var jitter time.Duration
	if e.MaxJitter > 0 {
		jitter = time.Duration(rand.Int63n(e.MaxJitter.Nanoseconds()))
	}
	if attempt < 0 {
		return e.Min + jitter
	}
	durFloat := float64(e.Min)
	durFloat += math.Pow(2, float64(attempt)) * float64(time.Second)
	dur := time.Duration(durFloat)
	if durFloat > float64(e.Max) {
		dur = e.Max
	}
	dur += jitter

	return dur
}

func Exponential() Strategy {
	return &ExponentialStrategy{
		Min:       0,
		Max:       10 * time.Second,
		MaxJitter: 250 * time.Millisecond,
	}
}

type FixedStrategy struct {
	Dur time.Duration
}

func (f *FixedStrategy) Duration(attempt int) time.Duration {
	return f.Dur
}

func Fixed(dur time.Duration) Strategy {
	return &FixedStrategy{
		Dur: dur,
	}
}
