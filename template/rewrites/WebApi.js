
                "url_rewrites": [
                  {
                    "path": "v4/breakingnews/{platform}?{key}",
                    "method": "GET",
                    "match_pattern": "v4\\/breakingnews\\/(.*)(\\?)(.*)",
                    "rewrite_to": "v4/breakingnews/$1?$3&global=false"
                  }
                ]
