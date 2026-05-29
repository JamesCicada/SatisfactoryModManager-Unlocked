package autoupdate

import (
	"github.com/satisfactorymodding/SatisfactoryModManager/backend/autoupdate/apply"
)

func init() {
	registerUpdateType("standalone", func() UpdateType {
		return UpdateType{
			ArtifactName: "SMMUnlocked_linux_amd64",
			Apply:        apply.MakeSingleFileApply(),
		}
	})
	registerUpdateType("appimage", func() UpdateType {
		return UpdateType{
			ArtifactName: "SMMUnlocked_linux_amd64.AppImage",
			Apply:        apply.MakeAppImageApply(),
		}
	})
}
