//ff:func feature=image type=parser control=sequence
//ff:what 변환 전/후 파일 크기를 조회해 ConvertResult로 조립
package img

import "os"

// statResult fills a ConvertResult with the on-disk sizes of src and dst.
func statResult(src, dst string) (ConvertResult, error) {
	si, err := os.Stat(src)
	if err != nil {
		return ConvertResult{}, err
	}
	di, err := os.Stat(dst)
	if err != nil {
		return ConvertResult{}, err
	}
	return ConvertResult{Dst: dst, SrcBytes: si.Size(), DstBytes: di.Size()}, nil
}
