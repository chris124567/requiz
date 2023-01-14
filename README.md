# requiz
Alternative Quizlet frontend in the style of Nitter/Invidious

# Motivation

My increasing frustration with Quizlet led me to develop this.  It is a substantially minified interface for Quizlet.  Simply replace the "quizlet.com" in any quizlet.com URL with requiz.net and you will get a minimalistic interface to Quizlet.  Supports learn/write mode for sets using Javascript.  Everything else is Javascript-free.

# How Does It Work

Emulates the mobile API and by default includes the API key of a throwaway account I made.  These keys never expire as far as I can tell.  Make sure you proxy (by setting HTTP_PROXY) through a residential IP address because api.quizlet.com now uses Cloudflare to block datacenter IPs.  Website is made in Bootstrap and the learn/write mode pages are written in Typescript with a library to neatly colorize string diffs for answers you get wrong.
