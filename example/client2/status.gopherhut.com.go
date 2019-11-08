package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	port := ":8082"
	http.HandleFunc("/", statusPageHandler)
	log.Println("status.gopherhut.com listening on ", port)
	http.ListenAndServe(port, nil)
}

func statusPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("incoming headers")
	for k, v := range r.Header {
		log.Printf("%s:%s\n", k, v)
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	tmpl := template.Must(template.New("status-page").Parse(statusPageHTML))
	tmpl.Execute(w, nil)
}

const statusPageHTML = `
<!DOCTYPE html>
<html lang="en" class="gr__githubstatus_com"><head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <!-- force IE browsers in compatibility mode to use their most aggressive rendering engine -->

    <meta charset="utf-8">
    <!-- <script type="text/javascript" src="https://bam.nr-data.net/1/bc12d0ca7c?a=1887052,5587075&amp;v=1130.54e767a&amp;to=IlgMRUFXWFhWEE5CQwBDF0IcS0BVRxcSHl4PUwdJ&amp;rst=2134&amp;ref=https://www.githubstatus.com/&amp;ap=95&amp;be=741&amp;fe=1615&amp;dc=1224&amp;perf=%7B%22timing%22:%7B%22of%22:1573136758913,%22n%22:0,%22f%22:12,%22dn%22:12,%22dne%22:12,%22c%22:12,%22ce%22:12,%22rq%22:27,%22rp%22:41,%22rpe%22:511,%22dl%22:59,%22di%22:1219,%22ds%22:1219,%22de%22:1260,%22dc%22:1615,%22l%22:1615,%22le%22:1616%7D,%22navigation%22:%7B%7D%7D&amp;fp=1262&amp;fcp=1262&amp;jsonp=NREUM.setToken"></script><script src="https://js-agent.newrelic.com/nr-1130.min.js"></script><script type="text/javascript" async="" src="https://www.gstatic.com/recaptcha/releases/0bBqi43w2fj-Lg1N3qzsqHNu/recaptcha__en.js"></script><script id="twitter-wjs" src="https://platform.twitter.com/widgets.js"></script><script type="text/javascript">window.NREUM||(NREUM={});NREUM.info={"beacon":"bam.nr-data.net","errorBeacon":"bam.nr-data.net","licenseKey":"bc12d0ca7c","applicationID":"1887052,5587075","transactionName":"IlgMRUFXWFhWEE5CQwBDF0IcS0BVRxcSHl4PUwdJ","queueTime":0,"applicationTime":95,"agent":""}</script> -->
    <script type="text/javascript">(window.NREUM||(NREUM={})).loader_config={licenseKey:"bc12d0ca7c",applicationID:"1887052"};window.NREUM||(NREUM={}),__nr_require=function(e,n,t){function r(t){if(!n[t]){var o=n[t]={exports:{}};e[t][0].call(o.exports,function(n){var o=e[t][1][n];return r(o||n)},o,o.exports)}return n[t].exports}if("function"==typeof __nr_require)return __nr_require;for(var o=0;o<t.length;o++)r(t[o]);return r}({1:[function(e,n,t){function r(){}function o(e,n,t){return function(){return i(e,[c.now()].concat(u(arguments)),n?null:this,t),n?void 0:this}}var i=e("handle"),a=e(3),u=e(4),f=e("ee").get("tracer"),c=e("loader"),s=NREUM;"undefined"==typeof window.newrelic&&(newrelic=s);var p=["setPageViewName","setCustomAttribute","setErrorHandler","finished","addToTrace","inlineHit","addRelease"],d="api-",l=d+"ixn-";a(p,function(e,n){s[n]=o(d+n,!0,"api")}),s.addPageAction=o(d+"addPageAction",!0),s.setCurrentRouteName=o(d+"routeName",!0),n.exports=newrelic,s.interaction=function(){return(new r).get()};var m=r.prototype={createTracer:function(e,n){var t={},r=this,o="function"==typeof n;return i(l+"tracer",[c.now(),e,t],r),function(){if(f.emit((o?"":"no-")+"fn-start",[c.now(),r,o],t),o)try{return n.apply(this,arguments)}catch(e){throw f.emit("fn-err",[arguments,this,e],t),e}finally{f.emit("fn-end",[c.now()],t)}}}};a("actionText,setName,setAttribute,save,ignore,onEnd,getContext,end,get".split(","),function(e,n){m[n]=o(l+n)}),newrelic.noticeError=function(e,n){"string"==typeof e&&(e=new Error(e)),i("err",[e,c.now(),!1,n])}},{}],2:[function(e,n,t){function r(e,n){if(!o)return!1;if(e!==o)return!1;if(!n)return!0;if(!i)return!1;for(var t=i.split("."),r=n.split("."),a=0;a<r.length;a++)if(r[a]!==t[a])return!1;return!0}var o=null,i=null,a=/Version\/(\S+)\s+Safari/;if(navigator.userAgent){var u=navigator.userAgent,f=u.match(a);f&&u.indexOf("Chrome")===-1&&u.indexOf("Chromium")===-1&&(o="Safari",i=f[1])}n.exports={agent:o,version:i,match:r}},{}],3:[function(e,n,t){function r(e,n){var t=[],r="",i=0;for(r in e)o.call(e,r)&&(t[i]=n(r,e[r]),i+=1);return t}var o=Object.prototype.hasOwnProperty;n.exports=r},{}],4:[function(e,n,t){function r(e,n,t){n||(n=0),"undefined"==typeof t&&(t=e?e.length:0);for(var r=-1,o=t-n||0,i=Array(o<0?0:o);++r<o;)i[r]=e[n+r];return i}n.exports=r},{}],5:[function(e,n,t){n.exports={exists:"undefined"!=typeof window.performance&&window.performance.timing&&"undefined"!=typeof window.performance.timing.navigationStart}},{}],ee:[function(e,n,t){function r(){}function o(e){function n(e){return e&&e instanceof r?e:e?f(e,u,i):i()}function t(t,r,o,i){if(!d.aborted||i){e&&e(t,r,o);for(var a=n(o),u=v(t),f=u.length,c=0;c<f;c++)u[c].apply(a,r);var p=s[y[t]];return p&&p.push([b,t,r,a]),a}}function l(e,n){h[e]=v(e).concat(n)}function m(e,n){var t=h[e];if(t)for(var r=0;r<t.length;r++)t[r]===n&&t.splice(r,1)}function v(e){return h[e]||[]}function g(e){return p[e]=p[e]||o(t)}function w(e,n){c(e,function(e,t){n=n||"feature",y[t]=n,n in s||(s[n]=[])})}var h={},y={},b={on:l,addEventListener:l,removeEventListener:m,emit:t,get:g,listeners:v,context:n,buffer:w,abort:a,aborted:!1};return b}function i(){return new r}function a(){(s.api||s.feature)&&(d.aborted=!0,s=d.backlog={})}var u="nr@context",f=e("gos"),c=e(3),s={},p={},d=n.exports=o();d.backlog=s},{}],gos:[function(e,n,t){function r(e,n,t){if(o.call(e,n))return e[n];var r=t();if(Object.defineProperty&&Object.keys)try{return Object.defineProperty(e,n,{value:r,writable:!0,enumerable:!1}),r}catch(i){}return e[n]=r,r}var o=Object.prototype.hasOwnProperty;n.exports=r},{}],handle:[function(e,n,t){function r(e,n,t,r){o.buffer([e],r),o.emit(e,n,t)}var o=e("ee").get("handle");n.exports=r,r.ee=o},{}],id:[function(e,n,t){function r(e){var n=typeof e;return!e||"object"!==n&&"function"!==n?-1:e===window?0:a(e,i,function(){return o++})}var o=1,i="nr@id",a=e("gos");n.exports=r},{}],loader:[function(e,n,t){function r(){if(!E++){var e=x.info=NREUM.info,n=l.getElementsByTagName("script")[0];if(setTimeout(s.abort,3e4),!(e&&e.licenseKey&&e.applicationID&&n))return s.abort();c(y,function(n,t){e[n]||(e[n]=t)}),f("mark",["onload",a()+x.offset],null,"api");var t=l.createElement("script");t.src="https://"+e.agent,n.parentNode.insertBefore(t,n)}}function o(){"complete"===l.readyState&&i()}function i(){f("mark",["domContent",a()+x.offset],null,"api")}function a(){return O.exists&&performance.now?Math.round(performance.now()):(u=Math.max((new Date).getTime(),u))-x.offset}var u=(new Date).getTime(),f=e("handle"),c=e(3),s=e("ee"),p=e(2),d=window,l=d.document,m="addEventListener",v="attachEvent",g=d.XMLHttpRequest,w=g&&g.prototype;NREUM.o={ST:setTimeout,SI:d.setImmediate,CT:clearTimeout,XHR:g,REQ:d.Request,EV:d.Event,PR:d.Promise,MO:d.MutationObserver};var h=""+location,y={beacon:"bam.nr-data.net",errorBeacon:"bam.nr-data.net",agent:"js-agent.newrelic.com/nr-1130.min.js"},b=g&&w&&w[m]&&!/CriOS/.test(navigator.userAgent),x=n.exports={offset:u,now:a,origin:h,features:{},xhrWrappable:b,userAgent:p};e(1),l[m]?(l[m]("DOMContentLoaded",i,!1),d[m]("load",r,!1)):(l[v]("onreadystatechange",o),d[v]("onload",r)),f("mark",["firstbyte",u],null,"api");var E=0,O=e(5)},{}]},{},["loader"]);</script>
    <title>GopherHut Status</title>
    <meta name="description" content="Welcome to GopherHut's home for real-time and historical data on system performance.">

    <!-- Mobile viewport optimization h5bp.com/ad -->
    <meta name="HandheldFriendly" content="True">
    <meta name="MobileOptimized" content="320">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0">

    <!-- Time this page was rendered - http://purl.org/dc/terms/issued -->
    <meta name="issued" content="1573136720">

    <!-- Mobile IE allows us to activate ClearType technology for smoothing fonts for easy reading -->
    <meta http-equiv="cleartype" content="on">

    <!-- Le fonts -->
<style>
  @font-face {
    font-family: 'proxima-nova';
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaLight-f0b2f7c12b6b87c65c02d3c1738047ea67a7607fd767056d8a2964cc6a2393f7.eot?host=www.githubstatus.com');
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaLight-f0b2f7c12b6b87c65c02d3c1738047ea67a7607fd767056d8a2964cc6a2393f7.eot?host=www.githubstatus.com#iefix') format('embedded-opentype'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaLight-e642ffe82005c6208632538a557e7f5dccb835c0303b06f17f55ccf567907241.woff?host=www.githubstatus.com') format('woff'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaLight-0f094da9b301d03292f97db5544142a16f9f2ddf50af91d44753d9310c194c5f.ttf?host=www.githubstatus.com') format('truetype');
    font-weight:300;
    font-style:normal;
  }
   
  @font-face {
    font-family: 'proxima-nova';
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaRegular-366d17769d864aa72f27defaddf591e460a1de4984bb24dacea57a9fc1d14878.eot?host=www.githubstatus.com');
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaRegular-366d17769d864aa72f27defaddf591e460a1de4984bb24dacea57a9fc1d14878.eot?host=www.githubstatus.com#iefix') format('embedded-opentype'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaRegular-2ee4c449a9ed716f1d88207bd1094e21b69e2818b5cd36b28ad809dc1924ec54.woff?host=www.githubstatus.com') format('woff'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaRegular-a40a469edbd27b65b845b8000d47445a17def8ba677f4eb836ad1808f7495173.ttf?host=www.githubstatus.com') format('truetype');
    font-weight:400;
    font-style:normal;
  }
   
  @font-face {
    font-family: 'proxima-nova';
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaRegularIt-0bf83a850b45e4ccda15bd04691e3c47ae84fec3588363b53618bd275a98cbb7.eot?host=www.githubstatus.com');
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaRegularIt-0bf83a850b45e4ccda15bd04691e3c47ae84fec3588363b53618bd275a98cbb7.eot?host=www.githubstatus.com#iefix') format('embedded-opentype'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaRegularIt-0c394ec7a111aa7928ea470ec0a67c44ebdaa0f93d1c3341abb69656cc26cbdd.woff?host=www.githubstatus.com') format('woff'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaRegularIt-9e43859f8015a4d47d9eaf7bafe8d1e26e3298795ce1f4cdb0be0479b8a4605e.ttf?host=www.githubstatus.com') format('truetype');
    font-weight:400;
    font-style:italic;
  }
   
  @font-face {
    font-family: 'proxima-nova';
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaSemibold-09566917307251d22021a3f91fc646f3e45f8d095209bcd2cded8a1979f06e54.eot?host=www.githubstatus.com');
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaSemibold-09566917307251d22021a3f91fc646f3e45f8d095209bcd2cded8a1979f06e54.eot?host=www.githubstatus.com#iefix') format('embedded-opentype'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaSemibold-86724fb2152613d735ba47c3f47a9ad2424b898bea4bece213dacee40344f966.woff?host=www.githubstatus.com') format('woff'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaSemibold-cf3e4eb7fbdf6fb83e526cc2a0141e55b01097e6e1abfd4cbdc3eda75d183f74.ttf?host=www.githubstatus.com') format('truetype');
    font-weight:500;
    font-style:normal;
  }
   
  @font-face {
    font-family: 'proxima-nova';
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaBold-622ea489d20e12e691663f83217105e957e2d3d09703707d40155a29c06cc9d9.eot?host=www.githubstatus.com');
    src: url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaBold-622ea489d20e12e691663f83217105e957e2d3d09703707d40155a29c06cc9d9.eot?host=www.githubstatus.com#iefix') format('embedded-opentype'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaBold-c8dc577ff7f76d2fc199843e38c04bb2e9fd15889421358d966a9f846c2ed1cd.woff?host=www.githubstatus.com') format('woff'),
         url('https://dka575ofm4ao0.cloudfront.net/assets/ProximaNovaBold-27177fe9242acbe089276ee587feef781446667ffe9b6fdc5b7fe21ad73e12f3.ttf?host=www.githubstatus.com') format('truetype');
    font-weight:700;
    font-style:normal;
  }
</style>


      <link rel="shortcut icon" type="image/x-icon" href="https://assets-cdn.github.com/favicon-success.ico">

    <link rel="shortcut icon" href="https://assets-cdn.github.com/favicon-success.ico">

    <link rel="alternate" type="application/atom+xml" href="https://www.githubstatus.com/history.atom" title="GitHub Status History - Atom Feed">
    <link rel="alternate" type="application/rss+xml" href="https://www.githubstatus.com/history.rss" title="GitHub Status History - RSS Feed">

      <!-- Canonical Link to ensure that only the custom domain is indexed when present -->
      <link rel="canonical" href="https://www.githubstatus.com">

    <meta name="_globalsign-domain-verification" content="y_VzfckMy4iePo5oDJNivyYIjh8LffYa4jzUndm_bZ">


    <link rel="alternate" type="application/atom+xml" title="ATOM" href="https://www.githubstatus.com/history.atom">

    <!-- Le styles -->
    <link rel="stylesheet" media="screen" href="https://dka575ofm4ao0.cloudfront.net/packs/0.74800ae6b852fb854747.css">
    <link rel="stylesheet" media="all" href="https://dka575ofm4ao0.cloudfront.net/assets/status/status_manifest-c1861740f6982c3c158247e83b19178001de400321368f741539965c6c184488.css">

    <script src="//ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
    <script>
      window.pageColorData = {"blue":"#0366d6","border":"#e1e4e8","body_background":"#ffffff","font":"#24292e","graph":"#0366d6","green":"#28a745","light_font":"#6a737d","link":"#0366d6","orange":"#e36209","red":"#dc3545","yellow":"#dbab09"};
    </script>
    <style>
  /* BODY BACKGROUND */ /* BODY BACKGROUND */ /* BODY BACKGROUND */ /* BODY BACKGROUND */ /* BODY BACKGROUND */
  body,
  .layout-content.status.status-api .section .example-container .example-opener .color-secondary,
  .grouped-items-selector,
  .layout-content.status.status-full-history .history-nav a.current,
  #uptime-tooltip .tooltip-box {
    background-color:#ffffff;
  }

  #uptime-tooltip .pointer-container .pointer-smaller {
    border-bottom-color:#ffffff;
  }

  /* PRIMARY FONT COLOR */ /* PRIMARY FONT COLOR */ /* PRIMARY FONT COLOR */ /* PRIMARY FONT COLOR */
  body.status,
  .color-primary,
  .color-primary:hover,
  .layout-content.status-index .status-day .update-title.impact-none a,
  .layout-content.status-index .status-day .update-title.impact-none a:hover,
  .layout-content.status-index .timeframes-container .timeframe.active,
  .layout-content.status-full-history .month .incident-container .impact-none,
  .layout-content.status.status-index .incidents-list .incident-title.impact-none a,
  .incident-history .impact-none,
  .layout-content.status .grouped-items-selector.inline .grouped-item.active,
  .layout-content.status.status-full-history .history-nav a.current,
  .layout-content.status.status-full-history .history-nav a:not(.current):hover,
  #uptime-tooltip .tooltip-box .tooltip-content .related-events .related-event a.related-event-link {
    color:#24292e;
  }

  .layout-content.status.status-index .components-statuses .component-container .name {
    color:#24292e;
    color:rgba(36,41,46,.8);
  }

  /* SECONDARY FONT COLOR */ /* SECONDARY FONT COLOR */ /* SECONDARY FONT COLOR */ /* SECONDARY FONT COLOR */
  small,
  .layout-content.status .table-row .date,
  .color-secondary,
  .layout-content.status .grouped-items-selector.inline .grouped-item,
  .layout-content.status.status-full-history .history-footer .pagination a.disabled,
  .layout-content.status.status-full-history .history-nav a,
  #uptime-tooltip .tooltip-box .tooltip-content .related-events #related-event-header {
    color:#6a737d;
  }

  /* BORDER COLOR */  /* BORDER COLOR */  /* BORDER COLOR */  /* BORDER COLOR */  /* BORDER COLOR */  /* BORDER COLOR */
  body.status .layout-content.status .border-color,
  hr,
  .tooltip-base,
  .markdown-display table,
  #uptime-tooltip .tooltip-box {
    border-color:#e1e4e8;
  }

  .markdown-display table td {
    border-top-color:#e1e4e8;
  }

  .markdown-display table td + td, .markdown-display table th + th {
    border-left-color:#e1e4e8;
  }

  #uptime-tooltip .pointer-container .pointer-larger {
    border-bottom-color:#e1e4e8;
  }

  #uptime-tooltip .tooltip-box .outage-field {
    background-color: rgba(225,228,232,0.31);
    border-color: rgba(225,228,232,0.0);
  }
  /* CSS REDS */ /* CSS REDS */ /* CSS REDS */ /* CSS REDS */ /* CSS REDS */ /* CSS REDS */ /* CSS REDS */
  .layout-content.status.status-index .status-day .update-title.impact-critical a,
  .layout-content.status.status-index .status-day .update-title.impact-critical a:hover,
  .layout-content.status.status-index .page-status.status-critical,
  .layout-content.status.status-index .unresolved-incident.impact-critical .incident-title,
  .flat-button.background-red {
    background-color:#dc3545;
  }

  .layout-content.status-index .components-statuses .component-container.status-red:after,
  .layout-content.status-full-history .month .incident-container .impact-critical,
  .layout-content.status-incident .incident-name.impact-critical,
  .layout-content.status.status-index .incidents-list .incident-title.impact-critical a,
  .status-red .icon-indicator,
  .incident-history .impact-critical,
  .components-container .component-inner-container.status-red .component-status,
  .components-container .component-inner-container.status-red .icon-indicator {
    color:#dc3545;
  }

  .layout-content.status.status-index .unresolved-incident.impact-critical .updates {
    border-color:#dc3545;
  }

  /* CSS ORANGES */ /* CSS ORANGES */ /* CSS ORANGES */ /* CSS ORANGES */ /* CSS ORANGES */ /* CSS ORANGES */
  .layout-content.status.status-index .status-day .update-title.impact-major a,
  .layout-content.status.status-index .status-day .update-title.impact-major a:hover,
  .layout-content.status.status-index .page-status.status-major,
  .layout-content.status.status-index .unresolved-incident.impact-major .incident-title {
    background-color:#e36209;
  }

  .layout-content.status-index .components-statuses .component-container.status-orange:after,
  .layout-content.status-full-history .month .incident-container .impact-major,
  .layout-content.status-incident .incident-name.impact-major,
  .layout-content.status.status-index .incidents-list .incident-title.impact-major a,
  .status-orange .icon-indicator,
  .incident-history .impact-major,
  .components-container .component-inner-container.status-orange .component-status,
  .components-container .component-inner-container.status-orange .icon-indicator {
    color:#e36209;
  }

  .layout-content.status.status-index .unresolved-incident.impact-major .updates {
    border-color:#e36209;
  }

  /* CSS YELLOWS */ /* CSS YELLOWS */ /* CSS YELLOWS */ /* CSS YELLOWS */ /* CSS YELLOWS */ /* CSS YELLOWS */
  .layout-content.status.status-index .status-day .update-title.impact-minor a,
  .layout-content.status.status-index .status-day .update-title.impact-minor a:hover,
  .layout-content.status.status-index .page-status.status-minor,
  .layout-content.status.status-index .unresolved-incident.impact-minor .incident-title,
  .layout-content.status.status-index .scheduled-incidents-container .tab {
    background-color:#dbab09;
  }

  .layout-content.status-index .components-statuses .component-container.status-yellow:after,
  .layout-content.status-full-history .month .incident-container .impact-minor,
  .layout-content.status-incident .incident-name.impact-minor,
  .layout-content.status.status-index .incidents-list .incident-title.impact-minor a,
  .status-yellow .icon-indicator,
  .incident-history .impact-minor,
  .components-container .component-inner-container.status-yellow .component-status,
  .components-container .component-inner-container.status-yellow .icon-indicator,
  .layout-content.status.manage-subscriptions .confirmation-infobox .fa {
    color:#dbab09;
  }

  .layout-content.status.status-index .unresolved-incident.impact-minor .updates,
  .layout-content.status.status-index .scheduled-incidents-container {
    border-color:#dbab09;
  }

  /* CSS BLUES */ /* CSS BLUES */ /* CSS BLUES */ /* CSS BLUES */ /* CSS BLUES */ /* CSS BLUES */
  .layout-content.status.status-index .status-day .update-title.impact-maintenance a,
  .layout-content.status.status-index .status-day .update-title.impact-maintenance a:hover,
  .layout-content.status.status-index .page-status.status-maintenance,
  .layout-content.status.status-index .unresolved-incident.impact-maintenance .incident-title,
  .layout-content.status.status-index .scheduled-incidents-container .tab {
    background-color:#0366d6;
  }

  .layout-content.status-index .components-statuses .component-container.status-blue:after,
  .layout-content.status-full-history .month .incident-container .impact-maintenance,
  .layout-content.status-incident .incident-name.impact-maintenance,
  .layout-content.status.status-index .incidents-list .incident-title.impact-maintenance a,
  .status-blue .icon-indicator,
  .incident-history .impact-maintenance,
  .components-container .component-inner-container.status-blue .component-status,
  .components-container .component-inner-container.status-blue .icon-indicator {
    color:#0366d6;
  }

  .layout-content.status.status-index .unresolved-incident.impact-maintenance .updates,
  .layout-content.status.status-index .scheduled-incidents-container {
    border-color:#0366d6;
  }

  /* CSS GREENS */ /* CSS GREENS */ /* CSS GREENS */ /* CSS GREENS */ /* CSS GREENS */ /* CSS GREENS */ /* CSS GREENS */
  .layout-content.status.status-index .page-status.status-none {
    background-color:#28a745;
  }
  .layout-content.status-index .components-statuses .component-container.status-green:after,
  .status-green .icon-indicator,
  .components-container .component-inner-container.status-green .component-status,
  .components-container .component-inner-container.status-green .icon-indicator {
    color:#28a745;
  }

  /* CSS LINK COLOR */  /* CSS LINK COLOR */  /* CSS LINK COLOR */  /* CSS LINK COLOR */  /* CSS LINK COLOR */  /* CSS LINK COLOR */
  a,
  a:hover,
  .layout-content.status-index .page-footer span a:hover,
  .layout-content.status-index .timeframes-container .timeframe:not(.active):hover,
  .layout-content.status-incident .subheader a:hover {
    color:#0366d6;
  }

  .flat-button,
  .masthead .updates-dropdown-container .show-updates-dropdown,
  .layout-content.status-full-history .show-filter.open  {
    background-color:#0366d6;
  }

  /* CUSTOM COLOR OVERRIDES FOR UPTIME SHOWCASE */
  .components-section .components-uptime-link {
    color: #6a737d;
  }

  .layout-content.status .shared-partial.uptime-90-days-wrapper .legend .legend-item {
    color: #6a737d;
    opacity: 0.8;
  }
  .layout-content.status .shared-partial.uptime-90-days-wrapper .legend .legend-item.light {
    color: #6a737d;
    opacity: 0.5;
  }
  .layout-content.status .shared-partial.uptime-90-days-wrapper .legend .spacer {
    background: #6a737d;
    opacity: 0.3;
  }
</style>


    <!-- custom css -->
    <link rel="stylesheet" type="text/css" href="//dka575ofm4ao0.cloudfront.net/page_display_customizations-custom_css_externals/36313/external20190830-55-w18ewe.css">
    <script charset="utf-8" src="https://platform.twitter.com/js/button.d941c9a422e2e3faf474b82a1f39e936.js"></script></head>
    <body class="status index status-none" data-gr-c-s-loaded="true" data-breakpoint-reached="false">

        <div class="layout-content status status-index starter">
            <div class="custom-header-container">
            <img src="https://user-images.githubusercontent.com/19292210/60553863-044dd200-9cea-11e9-987e-7db84449f215.png" class="illo-desktop-header">
        </div>


    <div class="container">
        <div class="page-status status-none">
          <span class="status font-large">
            GopherHut - All Systems Operational
          </span>
          <span class="last-updated-stamp  font-small"></span>
        </div>



        <div class="components-section font-regular">
    <div class="components-container one-column">
            <div class="component-container border-color">
              
<div data-component-id="8l4ygp009s5s" class="component-inner-container status-green " data-component-status="operational" data-js-hook="">

   <span class="name">
      Git Operations
   </span>

    <span class="tooltip-base tool tooltipstered">?</span>

  <span class="component-status " title="">Operational</span>

  <span class="tool icon-indicator fa fa-check tooltipstered"></span>

</div>

            </div>
            <div class="component-container border-color">
              
<div data-component-id="brv1bkgrwx7q" class="component-inner-container status-green " data-component-status="operational" data-js-hook="">

   <span class="name">
      API Requests
   </span>

    <span class="tooltip-base tool tooltipstered">?</span>

  <span class="component-status " title="">Operational</span>

  <span class="tool icon-indicator fa fa-check tooltipstered"></span>

</div>

            </div>
            <div class="component-container border-color">
              
<div data-component-id="kr09ddfgbfsf" class="component-inner-container status-green " data-component-status="operational" data-js-hook="">

   <span class="name">
      Issues, PRs, Dashboard, Projects
   </span>

    <span class="tooltip-base tool tooltipstered">?</span>

  <span class="component-status " title="">Operational</span>

  <span class="tool icon-indicator fa fa-check tooltipstered"></span>

</div>

            </div>
            <div class="component-container border-color" style="display: none;">
              
<div data-component-id="0l2p9nhqnxpd" class="component-inner-container status-green " data-component-status="operational" data-js-hook="">

   <span class="name">
      Visit www.githubstatus.com for more information
   </span>


  <span class="component-status " title="">Operational</span>

  <span class="tool icon-indicator fa fa-check tooltipstered"></span>

</div>

            </div>
            <div class="component-container border-color">
              
<div data-component-id="5bfcr2x9x8kc" class="component-inner-container status-green " data-component-status="operational" data-js-hook="">

   <span class="name">
      Notifications
   </span>

    <span class="tooltip-base tool tooltipstered">?</span>

  <span class="component-status " title="">Operational</span>

  <span class="tool icon-indicator fa fa-check tooltipstered"></span>

</div>

            </div>
            <div class="component-container border-color">
              
<div data-component-id="04c28ykz2c5m" class="component-inner-container status-green " data-component-status="operational" data-js-hook="">

   <span class="name">
      Gists
   </span>

    <span class="tooltip-base tool tooltipstered">?</span>

  <span class="component-status " title="">Operational</span>

  <span class="tool icon-indicator fa fa-check tooltipstered"></span>

</div>

            </div>
            <div class="component-container border-color">
              
<div data-component-id="vg70hn9s2tyj" class="component-inner-container status-green " data-component-status="operational" data-js-hook="">

   <span class="name">
      GitHub Pages
   </span>

    <span class="tooltip-base tool tooltipstered">?</span>

  <span class="component-status " title="">Operational</span>

  <span class="tool icon-indicator fa fa-check tooltipstered"></span>

</div>

            </div>
    </div>
    <div class="component-statuses-legend font-small">
  <div class="legend-item status-green">
    <span class="icon-indicator fa fa-check"></span>
    Operational
  </div>
  <div class="legend-item status-yellow">
    <span class="icon-indicator fa fa-minus-square"></span>
    Degraded Performance
  </div>
  <div class="legend-item status-orange">
    <span class="icon-indicator fa fa-exclamation-triangle"></span>
    Partial Outage
  </div>
  <div class="breaker"></div>
  <div class="legend-item status-red">
    <span class="icon-indicator fa fa-times"></span>
    Major Outage
  </div>
  <div class="legend-item status-blue">
    <span class="icon-indicator fa fa-wrench"></span>
    Maintenance
  </div>
</div>

  </div>






      <div class="incidents-list format-expanded">
        <a class="font-largest no-link" id="past-incidents" href="#past-incidents">Past Incidents</a>
          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Nov <var data-var="date"> 7</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported today.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Nov <var data-var="date"> 6</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular ">
    <div class="date border-color font-large">Nov <var data-var="date"> 5</var>, <var data-var="year">2019</var></div>
          <div class="incident-container">
  <div class="incident-title impact-minor font-large">
    <a href="/incidents/42hkbtl63nmn">Incident on 2019-11-05 13:42 UTC</a>
  </div>

  <div class="updates-container">
    <!-- postmortem -->

    <!-- incident updates -->
      <div class="update font-regular resolved">
        <strong>Resolved</strong> -
      	This incident has been resolved.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">20:44</var> UTC
        </small>
      </div>
      <div class="update font-regular update">
        <strong>Update</strong> -
      	We are continuing to monitor as the backlog clears. Users may see delays in search results.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">20:05</var> UTC
        </small>
      </div>
      <div class="update font-regular update">
        <strong>Update</strong> -
      	We are back to working through our backlog. Users should see webhook delivery delays as we recover.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">18:26</var> UTC
        </small>
      </div>
      <div class="update font-regular update">
        <strong>Update</strong> -
      	Despite our mitigation efforts, we are now experiencing a service interruption.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">18:09</var> UTC
        </small>
      </div>
      <div class="update font-regular update">
        <strong>Update</strong> -
      	Engineers have identified contributing factors to the backlog and have begun their mitigation work.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">17:38</var> UTC
        </small>
      </div>
      <div class="update font-regular update">
        <strong>Update</strong> -
      	We are continuing to investigate and closely monitor queue sizes.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">16:36</var> UTC
        </small>
      </div>
      <div class="update font-regular update">
        <strong>Update</strong> -
      	We are continuing to investigate higher than normal queue backlogs.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">15:31</var> UTC
        </small>
      </div>
      <div class="update font-regular update">
        <strong>Update</strong> -
      	We are continuing to monitor higher than normal queue backlogs.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">14:44</var> UTC
        </small>
      </div>
      <div class="update font-regular update">
        <strong>Update</strong> -
      	Queue backlogs are currently also causing delays in merging of PRs. We continue to investigate.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">13:48</var> UTC
        </small>
      </div>
      <div class="update font-regular investigating">
        <strong>Investigating</strong> -
      	Currently investigating increase in queue backlogs affecting Pages builds and some webhooks deliveries.

        <br>

        <small>
            Nov <var data-var="date"> 5</var>, <var data-var="time">13:42</var> UTC
        </small>
      </div>
  </div>

</div>

  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Nov <var data-var="date"> 4</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Nov <var data-var="date"> 3</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Nov <var data-var="date"> 2</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Nov <var data-var="date"> 1</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Oct <var data-var="date">31</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Oct <var data-var="date">30</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Oct <var data-var="date">29</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Oct <var data-var="date">28</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Oct <var data-var="date">27</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Oct <var data-var="date">26</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Oct <var data-var="date">25</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

          
  <div class="status-day font-regular no-incidents">
    <div class="date border-color font-large">Oct <var data-var="date">24</var>, <var data-var="year">2019</var></div>
        <p class="color-secondary">No incidents reported.</p>
  </div>

      <div class="custom-footer-container">
    <div class="footer mt-6 border-top" role="contentinfo">
  <img src="https://user-images.githubusercontent.com/19292210/60553864-044dd200-9cea-11e9-996a-a7a316ec3a35.png" class="illo-mobile-footer">
  
    </div>
  </div>
  </div>

<div id="cpt-notification-container"></div>
`
