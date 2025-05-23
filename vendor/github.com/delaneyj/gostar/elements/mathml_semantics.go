// DO NOT EDIT THIS FILE. IT IS AUTOMATICALLY GENERATED.
// package mathml semantics is generated from configuration file.
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

// This element is used to specify the meaning of a MathML expression.
type MathMLSEMANTICSElement struct {
	*Element
}

// Create a new MathMLSEMANTICSElement element.
// This will create a new element with the tag
// "semantics" during rendering.
func MathML_SEMANTICS(children ...ElementRenderer) *MathMLSEMANTICSElement {
	e := NewElement("semantics", children...)
	e.IsSelfClosing = false
	e.Descendants = children

	return &MathMLSEMANTICSElement{Element: e}
}

func (e *MathMLSEMANTICSElement) Children(children ...ElementRenderer) *MathMLSEMANTICSElement {
	e.Descendants = append(e.Descendants, children...)
	return e
}

func (e *MathMLSEMANTICSElement) IfChildren(condition bool, children ...ElementRenderer) *MathMLSEMANTICSElement {
	if condition {
		e.Descendants = append(e.Descendants, children...)
	}
	return e
}

func (e *MathMLSEMANTICSElement) TernChildren(condition bool, trueChildren, falseChildren ElementRenderer) *MathMLSEMANTICSElement {
	if condition {
		e.Descendants = append(e.Descendants, trueChildren)
	} else {
		e.Descendants = append(e.Descendants, falseChildren)
	}
	return e
}

func (e *MathMLSEMANTICSElement) Attr(name string, value string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set(name, value)
	return e
}

func (e *MathMLSEMANTICSElement) Attrs(attrs ...string) *MathMLSEMANTICSElement {
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

func (e *MathMLSEMANTICSElement) AttrsMap(attrs map[string]string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	for k, v := range attrs {
		e.StringAttributes.Set(k, v)
	}
	return e
}

func (e *MathMLSEMANTICSElement) Text(text string) *MathMLSEMANTICSElement {
	e.Descendants = append(e.Descendants, Text(text))
	return e
}

func (e *MathMLSEMANTICSElement) TextF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.Text(fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfText(condition bool, text string) *MathMLSEMANTICSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(text))
	}
	return e
}

func (e *MathMLSEMANTICSElement) IfTextF(condition bool, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.Descendants = append(e.Descendants, Text(fmt.Sprintf(format, args...)))
	}
	return e
}

func (e *MathMLSEMANTICSElement) Escaped(text string) *MathMLSEMANTICSElement {
	e.Descendants = append(e.Descendants, Escaped(text))
	return e
}

func (e *MathMLSEMANTICSElement) IfEscaped(condition bool, text string) *MathMLSEMANTICSElement {
	if condition {
		e.Descendants = append(e.Descendants, Escaped(text))
	}
	return e
}

func (e *MathMLSEMANTICSElement) EscapedF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.Escaped(fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfEscapedF(condition bool, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.Descendants = append(e.Descendants, EscapedF(format, args...))
	}
	return e
}

func (e *MathMLSEMANTICSElement) CustomData(key, value string) *MathMLSEMANTICSElement {
	if e.CustomDataAttributes == nil {
		e.CustomDataAttributes = treemap.New[string, string]()
	}
	e.CustomDataAttributes.Set(key, value)
	return e
}

func (e *MathMLSEMANTICSElement) IfCustomData(condition bool, key, value string) *MathMLSEMANTICSElement {
	if condition {
		e.CustomData(key, value)
	}
	return e
}

func (e *MathMLSEMANTICSElement) CustomDataF(key, format string, args ...any) *MathMLSEMANTICSElement {
	return e.CustomData(key, fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfCustomDataF(condition bool, key, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.CustomData(key, fmt.Sprintf(format, args...))
	}
	return e
}

func (e *MathMLSEMANTICSElement) CustomDataRemove(key string) *MathMLSEMANTICSElement {
	if e.CustomDataAttributes == nil {
		return e
	}
	e.CustomDataAttributes.Del(key)
	return e
}

// Assigns a class name or set of class names to an element
// You may assign the same class name or names to any number of elements
// If you specify multiple class names, they must be separated by whitespace
// characters.
func (e *MathMLSEMANTICSElement) CLASS(s ...string) *MathMLSEMANTICSElement {
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

func (e *MathMLSEMANTICSElement) IfCLASS(condition bool, s ...string) *MathMLSEMANTICSElement {
	if condition {
		e.CLASS(s...)
	}
	return e
}

// Remove the attribute CLASS from the element.
func (e *MathMLSEMANTICSElement) CLASSRemove(s ...string) *MathMLSEMANTICSElement {
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

// This attribute specifies the text directionality of the element, merely
// indicating what direction the text flows when surrounded by text with inherent
// directionality (such as Arabic or Hebrew)
// Possible values are ltr (left-to-right) and rtl (right-to-left).
func (e *MathMLSEMANTICSElement) DIR(c MathMLSemanticsDirChoice) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("dir", string(c))
	return e
}

type MathMLSemanticsDirChoice string

const (
	// left-to-right
	MathMLSemanticsDir_ltr MathMLSemanticsDirChoice = "ltr"
	// right-to-left
	MathMLSemanticsDir_rtl MathMLSemanticsDirChoice = "rtl"
)

// Remove the attribute DIR from the element.
func (e *MathMLSEMANTICSElement) DIRRemove(c MathMLSemanticsDirChoice) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("dir")
	return e
}

// This attribute specifies whether the element should be rendered using
// displaystyle rules or not
// Possible values are true and false.
func (e *MathMLSEMANTICSElement) DISPLAYSTYLE(c MathMLSemanticsDisplaystyleChoice) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("displaystyle", string(c))
	return e
}

type MathMLSemanticsDisplaystyleChoice string

const (
	// displaystyle rules
	MathMLSemanticsDisplaystyle_true MathMLSemanticsDisplaystyleChoice = "true"
	// not displaystyle rules
	MathMLSemanticsDisplaystyle_false MathMLSemanticsDisplaystyleChoice = "false"
)

// Remove the attribute DISPLAYSTYLE from the element.
func (e *MathMLSEMANTICSElement) DISPLAYSTYLERemove(c MathMLSemanticsDisplaystyleChoice) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("displaystyle")
	return e
}

// This attribute assigns a name to an element
// This name must be unique in a document.
func (e *MathMLSEMANTICSElement) ID(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("id", s)
	return e
}

func (e *MathMLSEMANTICSElement) IDF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.ID(fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfID(condition bool, s string) *MathMLSEMANTICSElement {
	if condition {
		e.ID(s)
	}
	return e
}

func (e *MathMLSEMANTICSElement) IfIDF(condition bool, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.ID(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute ID from the element.
func (e *MathMLSEMANTICSElement) IDRemove(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("id")
	return e
}

func (e *MathMLSEMANTICSElement) IDRemoveF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.IDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the background color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLSEMANTICSElement) MATHBACKGROUND(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathbackground", s)
	return e
}

func (e *MathMLSEMANTICSElement) MATHBACKGROUNDF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.MATHBACKGROUND(fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfMATHBACKGROUND(condition bool, s string) *MathMLSEMANTICSElement {
	if condition {
		e.MATHBACKGROUND(s)
	}
	return e
}

func (e *MathMLSEMANTICSElement) IfMATHBACKGROUNDF(condition bool, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.MATHBACKGROUND(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHBACKGROUND from the element.
func (e *MathMLSEMANTICSElement) MATHBACKGROUNDRemove(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathbackground")
	return e
}

func (e *MathMLSEMANTICSElement) MATHBACKGROUNDRemoveF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.MATHBACKGROUNDRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the color of the element
// Possible values are a color name or a color specification in the format defined
// in the CSS3 Color Module [CSS3COLOR].
func (e *MathMLSEMANTICSElement) MATHCOLOR(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathcolor", s)
	return e
}

func (e *MathMLSEMANTICSElement) MATHCOLORF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.MATHCOLOR(fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfMATHCOLOR(condition bool, s string) *MathMLSEMANTICSElement {
	if condition {
		e.MATHCOLOR(s)
	}
	return e
}

func (e *MathMLSEMANTICSElement) IfMATHCOLORF(condition bool, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.MATHCOLOR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHCOLOR from the element.
func (e *MathMLSEMANTICSElement) MATHCOLORRemove(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathcolor")
	return e
}

func (e *MathMLSEMANTICSElement) MATHCOLORRemoveF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.MATHCOLORRemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the size of the element
// Possible values are a dimension or a dimensionless number.
func (e *MathMLSEMANTICSElement) MATHSIZE_STR(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("mathsize", s)
	return e
}

func (e *MathMLSEMANTICSElement) MATHSIZE_STRF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.MATHSIZE_STR(fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfMATHSIZE_STR(condition bool, s string) *MathMLSEMANTICSElement {
	if condition {
		e.MATHSIZE_STR(s)
	}
	return e
}

func (e *MathMLSEMANTICSElement) IfMATHSIZE_STRF(condition bool, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.MATHSIZE_STR(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute MATHSIZE_STR from the element.
func (e *MathMLSEMANTICSElement) MATHSIZE_STRRemove(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("mathsize")
	return e
}

func (e *MathMLSEMANTICSElement) MATHSIZE_STRRemoveF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.MATHSIZE_STRRemove(fmt.Sprintf(format, args...))
}

// This attribute declares a cryptographic nonce (number used once) that should be
// used by the server processing the element’s submission, and the resulting
// resource must be delivered with a Content-Security-Policy nonce attribute
// matching the value of the nonce attribute.
func (e *MathMLSEMANTICSElement) NONCE(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}
	e.StringAttributes.Set("nonce", s)
	return e
}

func (e *MathMLSEMANTICSElement) NONCEF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.NONCE(fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfNONCE(condition bool, s string) *MathMLSEMANTICSElement {
	if condition {
		e.NONCE(s)
	}
	return e
}

func (e *MathMLSEMANTICSElement) IfNONCEF(condition bool, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.NONCE(fmt.Sprintf(format, args...))
	}
	return e
}

// Remove the attribute NONCE from the element.
func (e *MathMLSEMANTICSElement) NONCERemove(s string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("nonce")
	return e
}

func (e *MathMLSEMANTICSElement) NONCERemoveF(format string, args ...any) *MathMLSEMANTICSElement {
	return e.NONCERemove(fmt.Sprintf(format, args...))
}

// This attribute specifies the script level of the element
// Possible values are an integer between 0 and 7, inclusive.
func (e *MathMLSEMANTICSElement) SCRIPTLEVEL(i int) *MathMLSEMANTICSElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("scriptlevel", i)
	return e
}

func (e *MathMLSEMANTICSElement) IfSCRIPTLEVEL(condition bool, i int) *MathMLSEMANTICSElement {
	if condition {
		e.SCRIPTLEVEL(i)
	}
	return e
}

// Remove the attribute SCRIPTLEVEL from the element.
func (e *MathMLSEMANTICSElement) SCRIPTLEVELRemove(i int) *MathMLSEMANTICSElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("scriptlevel")
	return e
}

// This attribute offers advisory information about the element for which it is
// set.
func (e *MathMLSEMANTICSElement) STYLEF(k string, format string, args ...any) *MathMLSEMANTICSElement {
	return e.STYLE(k, fmt.Sprintf(format, args...))
}

func (e *MathMLSEMANTICSElement) IfSTYLE(condition bool, k string, v string) *MathMLSEMANTICSElement {
	if condition {
		e.STYLE(k, v)
	}
	return e
}

func (e *MathMLSEMANTICSElement) STYLE(k string, v string) *MathMLSEMANTICSElement {
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

func (e *MathMLSEMANTICSElement) IfSTYLEF(condition bool, k string, format string, args ...any) *MathMLSEMANTICSElement {
	if condition {
		e.STYLE(k, fmt.Sprintf(format, args...))
	}
	return e
}

// Add the attributes in the map to the element.
func (e *MathMLSEMANTICSElement) STYLEMap(m map[string]string) *MathMLSEMANTICSElement {
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
func (e *MathMLSEMANTICSElement) STYLEPairs(pairs ...string) *MathMLSEMANTICSElement {
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

func (e *MathMLSEMANTICSElement) IfSTYLEPairs(condition bool, pairs ...string) *MathMLSEMANTICSElement {
	if condition {
		e.STYLEPairs(pairs...)
	}
	return e
}

// Remove the attribute STYLE from the element.
func (e *MathMLSEMANTICSElement) STYLERemove(keys ...string) *MathMLSEMANTICSElement {
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

// This attribute specifies the position of the current element in the tabbing
// order for the current document
// This value must be a number between 0 and 32767
// User agents should ignore leading zeros.
func (e *MathMLSEMANTICSElement) TABINDEX(i int) *MathMLSEMANTICSElement {
	if e.IntAttributes == nil {
		e.IntAttributes = treemap.New[string, int]()
	}
	e.IntAttributes.Set("tabindex", i)
	return e
}

func (e *MathMLSEMANTICSElement) IfTABINDEX(condition bool, i int) *MathMLSEMANTICSElement {
	if condition {
		e.TABINDEX(i)
	}
	return e
}

// Remove the attribute TABINDEX from the element.
func (e *MathMLSEMANTICSElement) TABINDEXRemove(i int) *MathMLSEMANTICSElement {
	if e.IntAttributes == nil {
		return e
	}
	e.IntAttributes.Del("tabindex")
	return e
}

// Merges the singleton store with the given object

func (e *MathMLSEMANTICSElement) DATASTAR_STORE(v any) *MathMLSEMANTICSElement {
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

func (e *MathMLSEMANTICSElement) DATASTAR_REF(expression string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-ref"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_REF(condition bool, expression string) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_REF(expression)
	}
	return e
}

// Remove the attribute DATASTAR_REF from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_REFRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-ref")
	return e
}

// Sets the value of the element

func (e *MathMLSEMANTICSElement) DATASTAR_BIND(key string, expression string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-bind-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_BIND(condition bool, key string, expression string) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_BIND(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_BIND from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_BINDRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-bind")
	return e
}

// Sets the value of the element

func (e *MathMLSEMANTICSElement) DATASTAR_MODEL(expression string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-model"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_MODEL(condition bool, expression string) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_MODEL(expression)
	}
	return e
}

// Remove the attribute DATASTAR_MODEL from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_MODELRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-model")
	return e
}

// Sets the textContent of the element

func (e *MathMLSEMANTICSElement) DATASTAR_TEXT(expression string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-text"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_TEXT(condition bool, expression string) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_TEXT(expression)
	}
	return e
}

// Remove the attribute DATASTAR_TEXT from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_TEXTRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-text")
	return e
}

// Sets the event handler of the element

type MathMLSemanticsOnMod customDataKeyModifier

// Debounces the event handler
func MathMLSemanticsOnModDebounce(
	d time.Duration,
) MathMLSemanticsOnMod {
	return func() string {
		return fmt.Sprintf("debounce_%dms", d.Milliseconds())
	}
}

// Throttles the event handler
func MathMLSemanticsOnModThrottle(
	d time.Duration,
) MathMLSemanticsOnMod {
	return func() string {
		return fmt.Sprintf("throttle_%dms", d.Milliseconds())
	}
}

func (e *MathMLSEMANTICSElement) DATASTAR_ON(key string, expression string, modifiers ...MathMLSemanticsOnMod) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-on-%s", key)

	customMods := lo.Map(modifiers, func(m MathMLSemanticsOnMod, i int) customDataKeyModifier {
		return customDataKeyModifier(m)
	})
	key = customDataKey(key, customMods...)
	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_ON(condition bool, key string, expression string, modifiers ...MathMLSemanticsOnMod) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_ON(key, expression, modifiers...)
	}
	return e
}

// Remove the attribute DATASTAR_ON from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_ONRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-on")
	return e
}

// Sets the focus of the element

func (e *MathMLSEMANTICSElement) DATASTAR_FOCUSSet(b bool) *MathMLSEMANTICSElement {
	key := "data-focus"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLSEMANTICSElement) DATASTAR_FOCUS() *MathMLSEMANTICSElement {
	return e.DATASTAR_FOCUSSet(true)
}

// Sets the header of for fetch requests

func (e *MathMLSEMANTICSElement) DATASTAR_HEADER(key string, expression string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key = fmt.Sprintf("data-header-%s", key)

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_HEADER(condition bool, key string, expression string) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_HEADER(key, expression)
	}
	return e
}

// Remove the attribute DATASTAR_HEADER from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_HEADERRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-header")
	return e
}

// Sets the indicator selector for fetch requests

func (e *MathMLSEMANTICSElement) DATASTAR_FETCH_INDICATOR(expression string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-fetch-indicator"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_FETCH_INDICATOR(condition bool, expression string) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_FETCH_INDICATOR(expression)
	}
	return e
}

// Remove the attribute DATASTAR_FETCH_INDICATOR from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_FETCH_INDICATORRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-fetch-indicator")
	return e
}

// Sets the visibility of the element

func (e *MathMLSEMANTICSElement) DATASTAR_SHOWSet(b bool) *MathMLSEMANTICSElement {
	key := "data-show"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLSEMANTICSElement) DATASTAR_SHOW() *MathMLSEMANTICSElement {
	return e.DATASTAR_SHOWSet(true)
}

// Triggers the callback when the element intersects the viewport

func (e *MathMLSEMANTICSElement) DATASTAR_INTERSECTS(expression string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-intersects"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_INTERSECTS(condition bool, expression string) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_INTERSECTS(expression)
	}
	return e
}

// Remove the attribute DATASTAR_INTERSECTS from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_INTERSECTSRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-intersects")
	return e
}

// Teleports the element to the given selector

func (e *MathMLSEMANTICSElement) DATASTAR_TELEPORTSet(b bool) *MathMLSEMANTICSElement {
	key := "data-teleport"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLSEMANTICSElement) DATASTAR_TELEPORT() *MathMLSEMANTICSElement {
	return e.DATASTAR_TELEPORTSet(true)
}

// Scrolls the element into view

func (e *MathMLSEMANTICSElement) DATASTAR_SCROLL_INTO_VIEWSet(b bool) *MathMLSEMANTICSElement {
	key := "data-scroll-into-view"
	e.BoolAttributes.Set(key, b)
	return e
}

func (e *MathMLSEMANTICSElement) DATASTAR_SCROLL_INTO_VIEW() *MathMLSEMANTICSElement {
	return e.DATASTAR_SCROLL_INTO_VIEWSet(true)
}

// Setup the ViewTransitionAPI for the element

func (e *MathMLSEMANTICSElement) DATASTAR_VIEW_TRANSITION(expression string) *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		e.StringAttributes = treemap.New[string, string]()
	}

	key := "data-view-transition"

	e.StringAttributes.Set(key, expression)
	return e
}

func (e *MathMLSEMANTICSElement) IfDATASTAR_VIEW_TRANSITION(condition bool, expression string) *MathMLSEMANTICSElement {
	if condition {
		e.DATASTAR_VIEW_TRANSITION(expression)
	}
	return e
}

// Remove the attribute DATASTAR_VIEW_TRANSITION from the element.
func (e *MathMLSEMANTICSElement) DATASTAR_VIEW_TRANSITIONRemove() *MathMLSEMANTICSElement {
	if e.StringAttributes == nil {
		return e
	}
	e.StringAttributes.Del("data-view-transition")
	return e
}
