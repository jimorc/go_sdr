module github.com/jimorc/go_sdr

go 1.22.0

require github.com/pothosware/go-soapy-sdr v0.7.4
require internal/soapy_logging v1.0.0
replace internal/soapy_logging => "./internal/soapy_logging"
