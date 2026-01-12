package processing

import (
	"context"

	"github.com/imgproxy/imgproxy/v3/imagedata"
	"github.com/imgproxy/imgproxy/v3/options"
	"github.com/imgproxy/imgproxy/v3/security"
	"github.com/imgproxy/imgproxy/v3/vips"
)

func urlWatermark(pctx *pipelineContext, img *vips.Image, po *options.ProcessingOptions, imgdata *imagedata.ImageData) error {
	if len(po.URLWatermarks) > 0 {
		for _, v := range po.URLWatermarks {
			if v.Enabled {
				var (
					imgData *imagedata.ImageData
					err     error
				)
				if imgData, err = imagedata.Download(context.Background(), v.ImageURL, "url_watermark", imagedata.DownloadOptions{Header: nil, CookieJar: nil}, security.DefaultOptions()); err != nil {
					return err
				}

				if err = applyWatermark(img, imgData, &v.WatermarkOptions, pctx.dprScale, 1); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
