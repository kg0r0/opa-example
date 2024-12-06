# Ref: https://www.openpolicyagent.org/docs/latest/policy-performance/
package example 

import rego.v1

default allow := false

allow if {
	some user
	input.method == "GET"
	input.path = ["accounts", user]
	input.user == user
}

allow if {
	input.method == "GET"
	input.path == ["accounts", "report"]
	roles[input.user][_] == "admin"
}

allow if {
	input.method == "POST"
	input.path == ["accounts"]
	roles[input.user][_] == "admin"
}

roles := {
	"bob": ["admin", "hr"],
	"alice": ["procurement"],
}