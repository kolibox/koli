package api

import (
	"fmt"
	"net/http"
	"regexp"

	platform "kolihub.io/koli/pkg/apis/v1alpha1"
	gitutil "kolihub.io/koli/pkg/git/util"	
)

var (
	releasesRegexp = regexp.MustCompile("/releases.+")
)

// Authenticate validates if the provided credentials are valid
func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Printf("URL: %#v", r.URL)
	_, jwtTokenString, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", "Basic")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Authentication required (token null)\n")
		return
	}
	u, err := gitutil.DecodeUserToken(jwtTokenString, h.cnf.PlatformClientSecret, h.cnf.Auth0.PlatformPubKey)
	if err != nil {
		w.Header().Set("WWW-Authenticate", "Basic")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Authentication required (%s)\n", err)
		return
	}
	h.user = u

	// A system token is only allowed to download releases,
	if u.Type == platform.SystemTokenType &&
		r.Method != "GET" &&
		!releasesRegexp.MatchString(r.URL.Path) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Access Denied! Not allowed to access the resource\n")
		return
	}
	next(w, r)
}
