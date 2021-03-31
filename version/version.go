package version

import _ "embed"

//Version shows what version SecretKube is at currently
//go:embed VERSION.txt
var Version string
