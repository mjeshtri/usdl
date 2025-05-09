// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg feTurbulence is generated from configuration file.
// Description:
package elements

import (
	"fmt"
	"html"
	"time"

	"github.com/goccy/go-json"
	"github.com/igrmk/treemap/v2"
	"github.com/samber/lo"
)

// The <feTurbulence> SVG filter primitive creates an image using the Perlin
// turbulence function
// It allows the synthesis of artificial textures like clouds or marble.
type SVGFETURBULENCEElement struct {
	*Element
}

// Create a new SVGFETURBULENCEElement element.
// This will create a new element with the tag
// "feTurbulence" during rendering.
func SVG_FETURBULENCE(children ...ElementRenderer) *SVGFETURBULENCEElement {
	e := NewElement("feTurbulence", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGFETURBULENCEElement{Element: e}
}

func (e *SVGFETURBULENCEElement) Children(children ...ElementRenderer) *SVGFETURBULENCEElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGFETURBULENCEElement) IfChildren(condition bool, children ...ElementRenderer) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGFETURBULENCEElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGFETURBULENCEElement) Attr(name string, value string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGFETURBULENCEElement) Attrs(attrs ...string) *SVGFETURBULENCEElement {
	if len(attrs)%2 != 0 {
		panic("attrs must be a multiple of 2")
	}
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for i := 0; i < len(attrs); i += 2 {
		k := attrs[i]
		v := attrs[i+1]
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFETURBULENCEElement) AttrsMap(attrs map[string]string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGFETURBULENCEElement) Text(text string) *SVGFETURBULENCEElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGFETURBULENCEElement) TextF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfText(condition bool, text string) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGFETURBULENCEElement) IfTextF(condition bool, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGFETURBULENCEElement) Escaped(text string) *SVGFETURBULENCEElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGFETURBULENCEElement) IfEscaped(condition bool, text string) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGFETURBULENCEElement) EscapedF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfEscapedF(condition bool, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGFETURBULENCEElement) CustomData(key, value string) *SVGFETURBULENCEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGFETURBULENCEElement) IfCustomData(condition bool, key, value string) *SVGFETURBULENCEElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGFETURBULENCEElement) CustomDataF(key, format string, args ...any) *SVGFETURBULENCEElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGFETURBULENCEElement) CustomDataRemove(key string) *SVGFETURBULENCEElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The baseFrequency attribute represent the base frequencies in the X and Y
// directions of the turbulence function.
func (e *SVGFETURBULENCEElement) BASE_FREQUENCY(s string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("baseFrequency", s)
	return e
}

func (e *SVGFETURBULENCEElement) BASE_FREQUENCYF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.BASE_FREQUENCY(fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfBASE_FREQUENCY(condition bool, s string) *SVGFETURBULENCEElement {
	if condition {
		e.BASE_FREQUENCY(s)
	}
	return e
}

func (e *SVGFETURBULENCEElement) IfBASE_FREQUENCYF(condition bool, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.BASE_FREQUENCY(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute BASE_FREQUENCY from the element.
func (e *SVGFETURBULENCEElement) BASE_FREQUENCYRemove(s string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("baseFrequency")
	return e
}

func (e *SVGFETURBULENCEElement) BASE_FREQUENCYRemoveF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.BASE_FREQUENCYRemove(fmt.Sprintf(format, args...))
}

// The numOctaves attribute indicates the number of octaves to be used by the
// noise function.
func (e *SVGFETURBULENCEElement) NUM_OCTAVES(f float64) *SVGFETURBULENCEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("numOctaves", f)
	return e
}

func (e *SVGFETURBULENCEElement) IfNUM_OCTAVES(condition bool, f float64) *SVGFETURBULENCEElement {
	if condition {
		e.NUM_OCTAVES(f)
	}
	return e
}

// The seed attribute indicates which number to use to seed the random number
// generator.
func (e *SVGFETURBULENCEElement) SEED(f float64) *SVGFETURBULENCEElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("seed", f)
	return e
}

func (e *SVGFETURBULENCEElement) IfSEED(condition bool, f float64) *SVGFETURBULENCEElement {
	if condition {
		e.SEED(f)
	}
	return e
}

// The stitchTiles attribute indicates how the Perlin noise function should be
// tiled
// It is ignored if type is not set to 'turbulence'.
func (e *SVGFETURBULENCEElement) STITCH_TILES(c SVGFeTurbulenceStitchTilesChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("stitchTiles", string(c))
	return e
}

type SVGFeTurbulenceStitchTilesChoice string

const (
	// The <feTurbulence> SVG filter primitive creates an image using the Perlin
	// turbulence function
	// It allows the synthesis of artificial textures like clouds or marble.
	SVGFeTurbulenceStitchTiles_noStitch SVGFeTurbulenceStitchTilesChoice = "noStitch"
	// The <feTurbulence> SVG filter primitive creates an image using the Perlin
	// turbulence function
	// It allows the synthesis of artificial textures like clouds or marble.
	SVGFeTurbulenceStitchTiles_stitch SVGFeTurbulenceStitchTilesChoice = "stitch"
)

// Remove the attribute STITCH_TILES from the element.
func (e *SVGFETURBULENCEElement) STITCH_TILESRemove(c SVGFeTurbulenceStitchTilesChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("stitchTiles")
	return e
}

// The type of turbulence function.
func (e *SVGFETURBULENCEElement) TYPE(c SVGFeTurbulenceTypeChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("type", string(c))
	return e
}

type SVGFeTurbulenceTypeChoice string

const (
	// The type of turbulence function.
	SVGFeTurbulenceType_fractalNoise SVGFeTurbulenceTypeChoice = "fractalNoise"
	// The type of turbulence function.
	SVGFeTurbulenceType_turbulence SVGFeTurbulenceTypeChoice = "turbulence"
)

// Remove the attribute TYPE from the element.
func (e *SVGFETURBULENCEElement) TYPERemove(c SVGFeTurbulenceTypeChoice) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("type")
	return e
}

// Specifies a unique id for an element
func (e *SVGFETURBULENCEElement) ID(s string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGFETURBULENCEElement) IDF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfID(condition bool, s string) *SVGFETURBULENCEElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGFETURBULENCEElement) IfIDF(condition bool, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGFETURBULENCEElement) IDRemove(s string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGFETURBULENCEElement) IDRemoveF(format string, args ...any) *SVGFETURBULENCEElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGFETURBULENCEElement) CLASS(s ...string) *SVGFETURBULENCEElement {
	if e.DelimitedStrings == nil {
		e.DelimitedStrings = treemap.New[string, *DelimitedBuilder[string]]()
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		ds = NewDelimitedBuilder[string](" ")
		e.DelimitedStrings.Set("class", ds)
	}
	ds.Add(s...)
	return e
}

func (e *SVGFETURBULENCEElement) IfCLASS(condition bool, s ...string) *SVGFETURBULENCEElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGFETURBULENCEElement) CLASSRemove(s ...string) *SVGFETURBULENCEElement {
	if e.DelimitedStrings == nil {
		return e
	}
	ds, ok := e.DelimitedStrings.Get("class")
	if !ok {
		return e
	}
	ds.Remove(s...)
	return e
}

// Specifies an inline CSS style for an element
func (e *SVGFETURBULENCEElement) STYLEF(k string, format string, args ...any) *SVGFETURBULENCEElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGFETURBULENCEElement) IfSTYLE(condition bool, k string, v string) *SVGFETURBULENCEElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGFETURBULENCEElement) STYLE(k string, v string) *SVGFETURBULENCEElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	kv.Add(k, v)
	return e
}

func (e *SVGFETURBULENCEElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGFETURBULENCEElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGFETURBULENCEElement) STYLEMap(m map[string]string) *SVGFETURBULENCEElement {
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}
	for k, v := range m {
		kv.Add(k, v)
	}
	return e
}

// Add pairs of attributes to the element.
func (e *SVGFETURBULENCEElement) STYLEPairs(pairs ...string) *SVGFETURBULENCEElement {
	if len(pairs)%2 != 0 {
		panic("Must have an even number of pairs")
	}
	if e.KVStrings == nil {
		e.KVStrings = treemap.New[string, *KVBuilder]()
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		kv = NewKVBuilder(":", ";")
		e.KVStrings.Set("style", kv)
	}

	for i := 0; i < len(pairs); i += 2 {
		kv.Add(pairs[i], pairs[i+1])
	}

	return e
}

func (e *SVGFETURBULENCEElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGFETURBULENCEElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGFETURBULENCEElement) STYLERemove(keys ...string) *SVGFETURBULENCEElement {
	if e.KVStrings == nil {
		return e
	}
	kv, ok := e.KVStrings.Get("style")
	if !ok {
		return e
	}
	for _, k := range keys {
		kv.Remove(k)
	}
	return e
}

// Merges the singleton store with the given object

func (e *SVGFETURBULENCEElement) DATASTAR_STORE(v any) *SVGFETURBULENCEElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	e.CustomDataAttributes.Set("store", html.EscapeString(string(b)))
	return e
}

// Sets the reference of the element

func (e *SVGFETURBULENCEElement) DATASTAR_REF(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_REF(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_REFRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGFETURBULENCEElement) DATASTAR_BIND(key string, expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_BINDRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGFETURBULENCEElement) DATASTAR_MODEL(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_MODELRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGFETURBULENCEElement) DATASTAR_TEXT(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_TEXTRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGFeTurbulenceOnMod customDataKeyModifier

// Debounces the event handler
func SVGFeTurbulenceOnModDebounce(
	d time.Duration,
) SVGFeTurbulenceOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGFeTurbulenceOnModThrottle(
	d time.Duration,
) SVGFeTurbulenceOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGFETURBULENCEElement) DATASTAR_ON(key string, expression string, modifiers ...SVGFeTurbulenceOnMod) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGFeTurbulenceOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGFeTurbulenceOnMod) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_ONRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGFETURBULENCEElement) DATASTAR_FOCUSSet(b bool) *SVGFETURBULENCEElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFETURBULENCEElement) DATASTAR_FOCUS() *SVGFETURBULENCEElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGFETURBULENCEElement) DATASTAR_HEADER(key string, expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_HEADERRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGFETURBULENCEElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_FETCH_INDICATORRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGFETURBULENCEElement) DATASTAR_SHOWSet(b bool) *SVGFETURBULENCEElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFETURBULENCEElement) DATASTAR_SHOW() *SVGFETURBULENCEElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGFETURBULENCEElement) DATASTAR_INTERSECTS(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_INTERSECTSRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGFETURBULENCEElement) DATASTAR_TELEPORTSet(b bool) *SVGFETURBULENCEElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFETURBULENCEElement) DATASTAR_TELEPORT() *SVGFETURBULENCEElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGFETURBULENCEElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGFETURBULENCEElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGFETURBULENCEElement) DATASTAR_SCROLL_INTO_VIEW() *SVGFETURBULENCEElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGFETURBULENCEElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGFETURBULENCEElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGFETURBULENCEElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGFETURBULENCEElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGFETURBULENCEElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
