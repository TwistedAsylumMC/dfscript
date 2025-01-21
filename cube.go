package dfscript

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/dop251/goja"
	"github.com/go-gl/mathgl/mgl64"
)

func (r *Runtime) setupCube() error {
	return newObject().
		// axis.go
		Const("Axis", map[string]cube.Axis{
			"X": cube.X,
			"Y": cube.Y,
			"Z": cube.Z,
		}).
		Method("axes", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(cube.Axes())
		}).
		// bbox.go
		Method("box", func(c goja.FunctionCall) goja.Value {
			if len(c.Arguments) < 6 {
				panic(fmt.Errorf("box expects 6 arguments"))
			}
			box := cube.Box(
				c.Argument(0).ToFloat(),
				c.Argument(1).ToFloat(),
				c.Argument(2).ToFloat(),
				c.Argument(3).ToFloat(),
				c.Argument(4).ToFloat(),
				c.Argument(5).ToFloat(),
			)
			return r.vm.ToValue(box)
		}).
		Method("anyIntersections", func(c goja.FunctionCall) goja.Value {
			var boxes []cube.BBox
			if v, ok := c.Argument(0).Export().([]any); ok {
				for i, b := range v {
					if box, ok := b.(cube.BBox); ok {
						boxes = append(boxes, box)
					} else {
						panic(r.vm.NewTypeError("argument 0 index %d is not a cube.BBox", i))
					}
				}
			} else {
				panic(r.vm.NewTypeError("argument 0 is not an array"))
			}
			search, ok := c.Argument(1).Export().(cube.BBox)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.BBox, got %T", c.Argument(1).Export()))
			}
			return r.vm.ToValue(cube.AnyIntersections(boxes, search))
		}).
		// direction.go
		Const("Direction", map[string]cube.Direction{
			"North": cube.North,
			"East":  cube.East,
			"South": cube.South,
			"West":  cube.West,
		}).
		Method("directions", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(cube.Directions())
		}).
		// face.go
		Const("Face", map[string]cube.Face{
			"Up":    cube.FaceUp,
			"Down":  cube.FaceDown,
			"North": cube.FaceNorth,
			"East":  cube.FaceEast,
			"South": cube.FaceSouth,
			"West":  cube.FaceWest,
		}).
		Method("faces", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(cube.Faces())
		}).
		Method("horizontalFaces", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(cube.HorizontalFaces())
		}).
		// orientation.go
		Method("orientation", func(c goja.FunctionCall) goja.Value {
			o := cube.Orientation(c.Argument(0).ToInteger())
			return r.vm.ToValue(o)
		}).
		Method("orientationFromYaw", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(cube.OrientationFromYaw(c.Argument(0).ToFloat()))
		}).
		// pos.go
		Method("pos", func(c goja.FunctionCall) goja.Value {
			if len(c.Arguments) < 3 {
				panic(fmt.Errorf("pos expects 3 arguments"))
			}
			pos := cube.Pos{
				int(c.Argument(0).ToInteger()),
				int(c.Argument(1).ToInteger()),
				int(c.Argument(2).ToInteger()),
			}
			return r.vm.ToValue(pos)
		}).
		Method("posFromVec3", func(c goja.FunctionCall) goja.Value {
			v, ok := c.Argument(0).Export().(mgl64.Vec3)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a mgl64.Vec3, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(cube.PosFromVec3(v))
		}).
		// range.go
		Method("range", func(c goja.FunctionCall) goja.Value {
			if len(c.Arguments) < 2 {
				panic(fmt.Errorf("range expects 2 arguments"))
			}
			ra := cube.Range{
				int(c.Argument(0).ToInteger()),
				int(c.Argument(1).ToInteger()),
			}
			return r.vm.ToValue(ra)
		}).
		// rotation.go
		Method("rotation", func(c goja.FunctionCall) goja.Value {
			if len(c.Arguments) < 2 {
				panic(fmt.Errorf("rotation expects 2 arguments"))
			}
			rot := cube.Rotation{
				c.Argument(0).ToFloat(),
				c.Argument(1).ToFloat(),
			}
			return r.vm.ToValue(rot)
		}).
		Apply(r.vm, "cube")
}
