// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package svg marker is generated from configuration file.
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

// The <marker> SVG element defines the graphics that is to be used for drawing
// arrowheads or polymarkers on a given <path>, <line>, <polyline> or <polygon>
// element.
type SVGMARKERElement struct {
	*Element
}

// Create a new SVGMARKERElement element.
// This will create a new element with the tag
// "marker" during rendering.
func SVG_MARKER(children ...ElementRenderer) *SVGMARKERElement {
	e := NewElement("marker", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &SVGMARKERElement{Element: e}
}

func (e *SVGMARKERElement) Children(children ...ElementRenderer) *SVGMARKERElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *SVGMARKERElement) IfChildren(condition bool, children ...ElementRenderer) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *SVGMARKERElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *SVGMARKERElement) Attr(name string, value string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *SVGMARKERElement) Attrs(attrs ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) AttrsMap(attrs map[string]string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *SVGMARKERElement) Text(text string) *SVGMARKERElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *SVGMARKERElement) TextF(format string, args ...any) *SVGMARKERElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfText(condition bool, text string) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *SVGMARKERElement) IfTextF(condition bool, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *SVGMARKERElement) Escaped(text string) *SVGMARKERElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *SVGMARKERElement) IfEscaped(condition bool, text string) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *SVGMARKERElement) EscapedF(format string, args ...any) *SVGMARKERElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfEscapedF(condition bool, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *SVGMARKERElement) CustomData(key, value string) *SVGMARKERElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *SVGMARKERElement) IfCustomData(condition bool, key, value string) *SVGMARKERElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *SVGMARKERElement) CustomDataF(key, format string, args ...any) *SVGMARKERElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfCustomDataF(condition bool, key, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *SVGMARKERElement) CustomDataRemove(key string) *SVGMARKERElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// The x-axis coordinate of the reference point which is to be aligned exactly at
// the marker position.
func (e *SVGMARKERElement) REF_X(f float64) *SVGMARKERElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("refX", f)
	return e
}

func (e *SVGMARKERElement) IfREF_X(condition bool, f float64) *SVGMARKERElement {
	if condition {
		e.REF_X(f)
	}
	return e
}

// The y-axis coordinate of the reference point which is to be aligned exactly at
// the marker position.
func (e *SVGMARKERElement) REF_Y(f float64) *SVGMARKERElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("refY", f)
	return e
}

func (e *SVGMARKERElement) IfREF_Y(condition bool, f float64) *SVGMARKERElement {
	if condition {
		e.REF_Y(f)
	}
	return e
}

// The width of the marker viewport.
func (e *SVGMARKERElement) MARKER_WIDTH(f float64) *SVGMARKERElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("markerWidth", f)
	return e
}

func (e *SVGMARKERElement) IfMARKER_WIDTH(condition bool, f float64) *SVGMARKERElement {
	if condition {
		e.MARKER_WIDTH(f)
	}
	return e
}

// The height of the marker viewport.
func (e *SVGMARKERElement) MARKER_HEIGHT(f float64) *SVGMARKERElement {
	if e.FloatAttributes == nil {
		e.FloatAttributes = treemap.New[string, float64]()
	}
	e.FloatAttributes.Set("markerHeight", f)
	return e
}

func (e *SVGMARKERElement) IfMARKER_HEIGHT(condition bool, f float64) *SVGMARKERElement {
	if condition {
		e.MARKER_HEIGHT(f)
	}
	return e
}

// The orientation of the marker relative to the shape it is attached to.
func (e *SVGMARKERElement) ORIENT(c SVGMarkerOrientChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("orient", string(c))
	return e
}

type SVGMarkerOrientChoice string

const (
	// The orientation of the marker relative to the shape it is attached to.
	SVGMarkerOrient_auto SVGMarkerOrientChoice = "auto"
	// The orientation of the marker relative to the shape it is attached to.
	SVGMarkerOrient_auto_start_reverse SVGMarkerOrientChoice = "auto-start-reverse"
	// The orientation of the marker relative to the shape it is attached to.
	SVGMarkerOrient_angle SVGMarkerOrientChoice = "angle"
)

// Remove the attribute ORIENT from the element.
func (e *SVGMARKERElement) ORIENTRemove(c SVGMarkerOrientChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("orient")
	return e
}

// The coordinate system for the various length values within the marker.
func (e *SVGMARKERElement) MARKER_UNITS(c SVGMarkerMarkerUnitsChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("markerUnits", string(c))
	return e
}

type SVGMarkerMarkerUnitsChoice string

const (
	// The coordinate system for the various length values within the marker.
	SVGMarkerMarkerUnits_userSpaceOnUse SVGMarkerMarkerUnitsChoice = "userSpaceOnUse"
	// The coordinate system for the various length values within the marker.
	SVGMarkerMarkerUnits_strokeWidth SVGMarkerMarkerUnitsChoice = "strokeWidth"
)

// Remove the attribute MARKER_UNITS from the element.
func (e *SVGMARKERElement) MARKER_UNITSRemove(c SVGMarkerMarkerUnitsChoice) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("markerUnits")
	return e
}

// The position and size of the marker viewport (the bounds of the marker).
func (e *SVGMARKERElement) VIEW_BOX(s string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("viewBox", s)
	return e
}

func (e *SVGMARKERElement) VIEW_BOXF(format string, args ...any) *SVGMARKERElement {
	return e.VIEW_BOX(fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfVIEW_BOX(condition bool, s string) *SVGMARKERElement {
	if condition {
		e.VIEW_BOX(s)
	}
	return e
}

func (e *SVGMARKERElement) IfVIEW_BOXF(condition bool, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.VIEW_BOX(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute VIEW_BOX from the element.
func (e *SVGMARKERElement) VIEW_BOXRemove(s string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("viewBox")
	return e
}

func (e *SVGMARKERElement) VIEW_BOXRemoveF(format string, args ...any) *SVGMARKERElement {
	return e.VIEW_BOXRemove(fmt.Sprintf(format, args...))
}

// Specifies a unique id for an element
func (e *SVGMARKERElement) ID(s string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *SVGMARKERElement) IDF(format string, args ...any) *SVGMARKERElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfID(condition bool, s string) *SVGMARKERElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *SVGMARKERElement) IfIDF(condition bool, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *SVGMARKERElement) IDRemove(s string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *SVGMARKERElement) IDRemoveF(format string, args ...any) *SVGMARKERElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// Specifies one or more classnames for an element (refers to a class in a style
// sheet)
func (e *SVGMARKERElement) CLASS(s ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) IfCLASS(condition bool, s ...string) *SVGMARKERElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *SVGMARKERElement) CLASSRemove(s ...string) *SVGMARKERElement {
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
func (e *SVGMARKERElement) STYLEF(k string, format string, args ...any) *SVGMARKERElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *SVGMARKERElement) IfSTYLE(condition bool, k string, v string) *SVGMARKERElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *SVGMARKERElement) STYLE(k string, v string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) IfSTYLEF(condition bool, k string, format string, args ...any) *SVGMARKERElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *SVGMARKERElement) STYLEMap(m map[string]string) *SVGMARKERElement {
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
func (e *SVGMARKERElement) STYLEPairs(pairs ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) IfSTYLEPairs(condition bool, pairs ...string) *SVGMARKERElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *SVGMARKERElement) STYLERemove(keys ...string) *SVGMARKERElement {
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

func (e *SVGMARKERElement) DATASTAR_STORE(v any) *SVGMARKERElement {
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

func (e *SVGMARKERElement) DATASTAR_REF(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_REF(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *SVGMARKERElement) DATASTAR_REFRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *SVGMARKERElement) DATASTAR_BIND(key string, expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_BIND(condition bool, key string, expression string) *SVGMARKERElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *SVGMARKERElement) DATASTAR_BINDRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *SVGMARKERElement) DATASTAR_MODEL(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_MODEL(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *SVGMARKERElement) DATASTAR_MODELRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *SVGMARKERElement) DATASTAR_TEXT(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_TEXT(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *SVGMARKERElement) DATASTAR_TEXTRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type SVGMarkerOnMod customDataKeyModifier

// Debounces the event handler
func SVGMarkerOnModDebounce(
	d time.Duration,
) SVGMarkerOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func SVGMarkerOnModThrottle(
	d time.Duration,
) SVGMarkerOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *SVGMARKERElement) DATASTAR_ON(key string, expression string, modifiers ...SVGMarkerOnMod) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m SVGMarkerOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...SVGMarkerOnMod) *SVGMARKERElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *SVGMARKERElement) DATASTAR_ONRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *SVGMARKERElement) DATASTAR_FOCUSSet(b bool) *SVGMARKERElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGMARKERElement) DATASTAR_FOCUS() *SVGMARKERElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *SVGMARKERElement) DATASTAR_HEADER(key string, expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *SVGMARKERElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *SVGMARKERElement) DATASTAR_HEADERRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *SVGMARKERElement) DATASTAR_FETCH_INDICATOR(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *SVGMARKERElement) DATASTAR_FETCH_INDICATORRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *SVGMARKERElement) DATASTAR_SHOWSet(b bool) *SVGMARKERElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGMARKERElement) DATASTAR_SHOW() *SVGMARKERElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *SVGMARKERElement) DATASTAR_INTERSECTS(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *SVGMARKERElement) DATASTAR_INTERSECTSRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *SVGMARKERElement) DATASTAR_TELEPORTSet(b bool) *SVGMARKERElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGMARKERElement) DATASTAR_TELEPORT() *SVGMARKERElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *SVGMARKERElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *SVGMARKERElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *SVGMARKERElement) DATASTAR_SCROLL_INTO_VIEW() *SVGMARKERElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *SVGMARKERElement) DATASTAR_VIEW_TRANSITION(expression string) *SVGMARKERElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *SVGMARKERElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *SVGMARKERElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *SVGMARKERElement) DATASTAR_VIEW_TRANSITIONRemove() *SVGMARKERElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
