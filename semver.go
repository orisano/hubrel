package main

import "github.com/blang/semver"

func nextMajor(v semver.Version) semver.Version {
	return semver.Version{
		Major: v.Major + 1,
		Minor: 0,
		Patch: 0,
	}
}

func nextMinor(v semver.Version) semver.Version {
	return semver.Version{
		Major: v.Major,
		Minor: v.Minor + 1,
		Patch: 0,
	}
}

func nextPatch(v semver.Version) semver.Version {
	return semver.Version{
		Major: v.Major,
		Minor: v.Minor,
		Patch: v.Patch + 1,
	}
}
