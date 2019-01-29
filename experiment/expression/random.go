package expression

import (
	"crypto/sha1"
	"fmt"
	"github.com/recursivecurry/golanout/experiment/base"
	"strconv"
	"strings"
)

const LongScale = float64(0xFFFFFFFFFFFFFFF)

type Random struct {
	Unit []base.Value
	Ctx *base.Context
}

func (r Random) getUnit(appendedUnit ...base.Value) []string {
	units := make([]string, 0, len(r.Unit)+len(appendedUnit))
	for _, u := range r.Unit {
		units = append(units, fmt.Sprintf("%+v", u))
	}
	for _, u := range appendedUnit {
		units = append(units, fmt.Sprintf("%+v", u))
	}
	return units
}

func (r Random) getHash(appendedUnit ...base.Value) int {
	fullSalt, ok := r.Ctx.Variables[base.VarFullSalt]
	if ok {
		fullSalt = fmt.Sprintf("%+v.", fullSalt)
	} else {
		salt, _ := r.Ctx.Variables[base.VarSalt]
		fullSalt = fmt.Sprintf("%s.%+v.", r.Ctx.Salt, salt)
	}

	unitString := strings.Join(r.getUnit(appendedUnit...), ".")
	hashString := fmt.Sprintf("%s%s", fullSalt, unitString)
	digest := sha1.Sum([]byte(hashString))
	hexDigest := fmt.Sprintf("%x", digest[:8])[:15]
	value, _ := strconv.ParseInt(hexDigest, 16, 64)
	return int(value)
}

func (r Random) getUniform(minValue, MaxValue float64, appendedUnit ...base.Value) float64 {
	zeroToOne := float64(r.getHash(appendedUnit...)) / LongScale
	return minValue + (MaxValue - minValue) * zeroToOne
}