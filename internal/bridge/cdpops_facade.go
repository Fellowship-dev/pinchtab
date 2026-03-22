package bridge

import (
	"context"
	"time"

	bridgecdpops "github.com/pinchtab/pinchtab/internal/bridge/cdpops"
)

const TargetTypePage = bridgecdpops.TargetTypePage

var (
	ImageBlockPatterns  = bridgecdpops.ImageBlockPatterns
	MediaBlockPatterns  = bridgecdpops.MediaBlockPatterns
	ErrTooManyRedirects = bridgecdpops.ErrTooManyRedirects
)

func NavigatePage(ctx context.Context, url string) error {
	return bridgecdpops.NavigatePage(ctx, url)
}

func NavigatePageWithRedirectLimit(ctx context.Context, url string, maxRedirects int) error {
	return bridgecdpops.NavigatePageWithRedirectLimit(ctx, url, maxRedirects)
}

func shouldReplaceBlankHistoryEntry(curURL string, cur int64, entryCount int) bool {
	return bridgecdpops.ShouldReplaceBlankHistoryEntry(curURL, cur, entryCount)
}

func WaitForTitle(ctx context.Context, timeout time.Duration) (string, error) {
	return bridgecdpops.WaitForTitle(ctx, timeout)
}

func SetResourceBlocking(ctx context.Context, patterns []string) error {
	return bridgecdpops.SetResourceBlocking(ctx, patterns)
}

func getElementCenterJS(ctx context.Context, backendNodeID int64) (float64, float64, error) {
	return bridgecdpops.GetElementCenterJS(ctx, backendNodeID)
}

func ScrollIntoViewAndGetBox(ctx context.Context, nodeID int64) (map[string]any, error) {
	return bridgecdpops.ScrollIntoViewAndGetBox(ctx, nodeID)
}

func ClickByCoordinate(ctx context.Context, x, y float64) error {
	return bridgecdpops.ClickByCoordinate(ctx, x, y)
}

func ClickByNodeID(ctx context.Context, nodeID int64) error {
	return bridgecdpops.ClickByNodeID(ctx, nodeID)
}

func DoubleClickByCoordinate(ctx context.Context, x, y float64) error {
	return bridgecdpops.DoubleClickByCoordinate(ctx, x, y)
}

func DoubleClickByNodeID(ctx context.Context, nodeID int64) error {
	return bridgecdpops.DoubleClickByNodeID(ctx, nodeID)
}

func DragByNodeID(ctx context.Context, nodeID int64, dx, dy int) error {
	return bridgecdpops.DragByNodeID(ctx, nodeID, dx, dy)
}

func DragByCoordinate(ctx context.Context, x, y float64, dx, dy int) error {
	return bridgecdpops.DragByCoordinate(ctx, x, y, dx, dy)
}

func HTML5DragByNodeID(ctx context.Context, sourceNodeID, targetNodeID int64) error {
	return bridgecdpops.HTML5DragByNodeID(ctx, sourceNodeID, targetNodeID)
}

func MouseDownByCoordinate(ctx context.Context, x, y float64) error {
	return bridgecdpops.MouseDownByCoordinate(ctx, x, y)
}

func MouseDownByNodeID(ctx context.Context, nodeID int64) error {
	return bridgecdpops.MouseDownByNodeID(ctx, nodeID)
}

func MouseUpByCoordinate(ctx context.Context, x, y float64) error {
	return bridgecdpops.MouseUpByCoordinate(ctx, x, y)
}

func MouseUpByNodeID(ctx context.Context, nodeID int64) error {
	return bridgecdpops.MouseUpByNodeID(ctx, nodeID)
}

func HoverByCoordinate(ctx context.Context, x, y float64) error {
	return bridgecdpops.HoverByCoordinate(ctx, x, y)
}

func HoverByNodeID(ctx context.Context, nodeID int64) error {
	return bridgecdpops.HoverByNodeID(ctx, nodeID)
}

func FillByNodeID(ctx context.Context, nodeID int64, value string) error {
	return bridgecdpops.FillByNodeID(ctx, nodeID, value)
}

func SelectByNodeID(ctx context.Context, nodeID int64, value string) error {
	return bridgecdpops.SelectByNodeID(ctx, nodeID, value)
}

func ScrollByNodeID(ctx context.Context, nodeID int64) error {
	return bridgecdpops.ScrollByNodeID(ctx, nodeID)
}
