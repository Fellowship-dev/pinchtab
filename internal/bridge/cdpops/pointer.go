package cdpops

import (
	"context"
	"fmt"
	"math"

	"github.com/chromedp/chromedp"
)

func ClickByCoordinate(ctx context.Context, x, y float64) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("x/y coordinates must be >= 0")
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mousePressed",
				"button":     "left",
				"clickCount": 1,
				"x":          x,
				"y":          y,
			}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mouseReleased",
				"button":     "left",
				"clickCount": 1,
				"x":          x,
				"y":          y,
			}, nil)
		}),
	)
}

func ClickByNodeID(ctx context.Context, nodeID int64) error {
	x, y, err := GetElementCenter(ctx, nodeID)
	if err != nil {
		return err
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.scrollIntoViewIfNeeded", map[string]any{"backendNodeId": nodeID}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.focus", map[string]any{"backendNodeId": nodeID}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mousePressed",
				"button":     "left",
				"clickCount": 1,
				"x":          x, "y": y,
			}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mouseReleased",
				"button":     "left",
				"clickCount": 1,
				"x":          x, "y": y,
			}, nil)
		}),
	)
}

func DoubleClickByCoordinate(ctx context.Context, x, y float64) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("x/y coordinates must be >= 0")
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mousePressed",
				"button":     "left",
				"clickCount": 2,
				"x":          x,
				"y":          y,
			}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mouseReleased",
				"button":     "left",
				"clickCount": 2,
				"x":          x,
				"y":          y,
			}, nil)
		}),
	)
}

func DoubleClickByNodeID(ctx context.Context, nodeID int64) error {
	x, y, err := GetElementCenter(ctx, nodeID)
	if err != nil {
		return err
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.scrollIntoViewIfNeeded", map[string]any{"backendNodeId": nodeID}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.focus", map[string]any{"backendNodeId": nodeID}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mousePressed",
				"button":     "left",
				"clickCount": 2,
				"x":          x, "y": y,
			}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mouseReleased",
				"button":     "left",
				"clickCount": 2,
				"x":          x, "y": y,
			}, nil)
		}),
	)
}

// DragByNodeID drags an element by (dx, dy) pixels using mousePressed → mouseMoved → mouseReleased.
func DragByNodeID(ctx context.Context, nodeID int64, dx, dy int) error {
	x, y, err := GetElementCenter(ctx, nodeID)
	if err != nil {
		return err
	}

	endX := x + float64(dx)
	endY := y + float64(dy)
	dist := math.Sqrt(float64(dx*dx + dy*dy))
	steps := int(dist / 10)
	if steps < 5 {
		steps = 5
	}
	if steps > 40 {
		steps = 40
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.scrollIntoViewIfNeeded", map[string]any{"backendNodeId": nodeID}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type": "mouseMoved",
				"x":    x, "y": y,
			}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mousePressed",
				"button":     "left",
				"clickCount": 1,
				"x":          x, "y": y,
			}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for i := 1; i <= steps; i++ {
				t := float64(i) / float64(steps)
				mx := x + t*float64(dx)
				my := y + t*float64(dy)
				if err := chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
					"type":    "mouseMoved",
					"buttons": 1,
					"x":       mx, "y": my,
				}, nil); err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mouseReleased",
				"button":     "left",
				"clickCount": 1,
				"x":          endX, "y": endY,
			}, nil)
		}),
	)
}

func MouseDownByCoordinate(ctx context.Context, x, y float64) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("x/y coordinates must be >= 0")
	}

	return chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
			"type":       "mousePressed",
			"button":     "left",
			"clickCount": 1,
			"x":          x,
			"y":          y,
		}, nil)
	}))
}

func MouseDownByNodeID(ctx context.Context, nodeID int64) error {
	x, y, err := GetElementCenter(ctx, nodeID)
	if err != nil {
		return err
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.scrollIntoViewIfNeeded", map[string]any{"backendNodeId": nodeID}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type":       "mousePressed",
				"button":     "left",
				"clickCount": 1,
				"x":          x, "y": y,
			}, nil)
		}),
	)
}

func MouseUpByCoordinate(ctx context.Context, x, y float64) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("x/y coordinates must be >= 0")
	}

	return chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
			"type":       "mouseReleased",
			"button":     "left",
			"clickCount": 1,
			"x":          x,
			"y":          y,
		}, nil)
	}))
}

func MouseUpByNodeID(ctx context.Context, nodeID int64) error {
	x, y, err := GetElementCenter(ctx, nodeID)
	if err != nil {
		return err
	}

	return chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
			"type":       "mouseReleased",
			"button":     "left",
			"clickCount": 1,
			"x":          x, "y": y,
		}, nil)
	}))
}

// HTML5DragByNodeID performs a synthetic HTML5 Drag and Drop between two elements.
// Uses JS to dispatch dragstart, dragover, drop, and dragend events.
func HTML5DragByNodeID(ctx context.Context, sourceNodeID, targetNodeID int64) error {
	js := `
(function(srcId, tgtId) {
  function nodeById(id) {
    return document.querySelector('[data-ptab-nid="' + id + '"]') ||
           (function() {
             const all = document.querySelectorAll('*');
             for (const el of all) if (el.__backendNodeId === id) return el;
             return null;
           })();
  }
  // Resolve via CDP-injected attribute or fallback
  const src = nodeById(srcId);
  const tgt = nodeById(tgtId);
  if (!src || !tgt) return JSON.stringify({error: 'element not found', src: !!src, tgt: !!tgt});

  const dt = new DataTransfer();
  const opts = {bubbles: true, cancelable: true, dataTransfer: dt};
  src.dispatchEvent(new DragEvent('dragstart', opts));
  tgt.dispatchEvent(new DragEvent('dragenter', opts));
  tgt.dispatchEvent(new DragEvent('dragover', opts));
  tgt.dispatchEvent(new DragEvent('drop', opts));
  src.dispatchEvent(new DragEvent('dragend', opts));
  return JSON.stringify({ok: true});
})
`
	// We need to resolve backendNodeId to DOM elements. Use DOM.resolveNode + evaluate.
	return chromedp.Run(ctx,
		// First, tag both elements with a data attribute so JS can find them
		chromedp.ActionFunc(func(ctx context.Context) error {
			js := fmt.Sprintf(`document.querySelector('[data-ptab-nid="%d"]')?.removeAttribute('data-ptab-nid'); void 0`, sourceNodeID)
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Runtime.evaluate", map[string]any{"expression": js}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Resolve source node and tag it
			var result map[string]any
			if err := chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.resolveNode", map[string]any{
				"backendNodeId": sourceNodeID,
			}, &result); err != nil {
				return fmt.Errorf("resolve source: %w", err)
			}
			objectID, _ := result["object"].(map[string]any)["objectId"].(string)
			if objectID == "" {
				return fmt.Errorf("could not resolve source node %d", sourceNodeID)
			}
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Runtime.callFunctionOn", map[string]any{
				"objectId":            objectID,
				"functionDeclaration": fmt.Sprintf(`function() { this.setAttribute('data-ptab-nid', '%d'); }`, sourceNodeID),
			}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Resolve target node and tag it
			var result map[string]any
			if err := chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.resolveNode", map[string]any{
				"backendNodeId": targetNodeID,
			}, &result); err != nil {
				return fmt.Errorf("resolve target: %w", err)
			}
			objectID, _ := result["object"].(map[string]any)["objectId"].(string)
			if objectID == "" {
				return fmt.Errorf("could not resolve target node %d", targetNodeID)
			}
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Runtime.callFunctionOn", map[string]any{
				"objectId":            objectID,
				"functionDeclaration": fmt.Sprintf(`function() { this.setAttribute('data-ptab-nid', '%d'); }`, targetNodeID),
			}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			expr := fmt.Sprintf(`%s(%d, %d)`, js, sourceNodeID, targetNodeID)
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Runtime.evaluate", map[string]any{"expression": expr}, nil)
		}),
		// Clean up tags
		chromedp.ActionFunc(func(ctx context.Context) error {
			cleanup := fmt.Sprintf(`document.querySelectorAll('[data-ptab-nid]').forEach(el => el.removeAttribute('data-ptab-nid')); void 0`)
			_ = cleanup
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Runtime.evaluate", map[string]any{
				"expression": `document.querySelectorAll('[data-ptab-nid]').forEach(el => el.removeAttribute('data-ptab-nid')); void 0`,
			}, nil)
		}),
	)
}

func HoverByCoordinate(ctx context.Context, x, y float64) error {
	if x < 0 || y < 0 {
		return fmt.Errorf("x/y coordinates must be >= 0")
	}

	return chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
			"type": "mouseMoved",
			"x":    x,
			"y":    y,
		}, nil)
	}))
}

func HoverByNodeID(ctx context.Context, nodeID int64) error {
	x, y, err := GetElementCenter(ctx, nodeID)
	if err != nil {
		return err
	}

	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "DOM.scrollIntoViewIfNeeded", map[string]any{"backendNodeId": nodeID}, nil)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.FromContext(ctx).Target.Execute(ctx, "Input.dispatchMouseEvent", map[string]any{
				"type": "mouseMoved",
				"x":    x, "y": y,
			}, nil)
		}),
	)
}
