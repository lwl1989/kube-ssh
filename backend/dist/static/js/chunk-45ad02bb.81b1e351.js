(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-45ad02bb"],{"4ec3":function(e,t,n){"use strict";n.d(t,"a",(function(){return c})),n.d(t,"e",(function(){return l})),n.d(t,"f",(function(){return d})),n.d(t,"g",(function(){return f})),n.d(t,"b",(function(){return p})),n.d(t,"d",(function(){return m})),n.d(t,"c",(function(){return h})),n.d(t,"k",(function(){return b})),n.d(t,"j",(function(){return v})),n.d(t,"i",(function(){return g})),n.d(t,"h",(function(){return w}));n("d3b7");var a=n("bc3a"),r=n.n(a),o=n("5c96"),s=n("5f87"),i=r.a.create({baseURL:"/api",timeout:5e3});i.interceptors.request.use((function(e){var t=Object(s["a"])();return t&&(e.headers.Signature=t),e}),(function(e){return console.log(e),Promise.reject(e)})),i.interceptors.response.use((function(e){var t=e.data;return 1e4!==t.code?(Object(o["Message"])({message:t.msg||"Error",type:"error",duration:5e3}),Promise.reject(new Error(t.msg||"Error"))):t}),(function(e){return console.log("err"+e),Object(o["Message"])({message:e.message,type:"error",duration:5e3}),Promise.reject(e)}));var u=i;function c(){return u({url:"/k8s/clusters",method:"get"})}function l(e){return u({url:"/k8s/workload/pods?id="+e.id,method:"get"})}function d(e){return u({url:"/k8s/sign",method:"post",data:e})}function f(){return u({url:"/users",method:"get"})}function p(e){return u({url:"/managers?page="+e.page+"&size="+e.size,method:"get"})}function m(e){return u({url:"/manager/upsert",method:"post",data:e})}function h(e){return u({url:"/manager/status",method:"post",data:e})}function b(e){return u({url:"/whites?page="+e.page+"&size="+e.size,method:"get"})}function v(e){return u({url:"/white",method:"post",data:e})}function g(e){return u({url:"/white/status",method:"post",data:e})}function w(e){return u({url:"/white",method:"delete",data:e})}},6724:function(e,t,n){"use strict";n("8d41");var a="@@wavesContext";function r(e,t){function n(n){var a=Object.assign({},t.value),r=Object.assign({ele:e,type:"hit",color:"rgba(0, 0, 0, 0.15)"},a),o=r.ele;if(o){o.style.position="relative",o.style.overflow="hidden";var s=o.getBoundingClientRect(),i=o.querySelector(".waves-ripple");switch(i?i.className="waves-ripple":(i=document.createElement("span"),i.className="waves-ripple",i.style.height=i.style.width=Math.max(s.width,s.height)+"px",o.appendChild(i)),r.type){case"center":i.style.top=s.height/2-i.offsetHeight/2+"px",i.style.left=s.width/2-i.offsetWidth/2+"px";break;default:i.style.top=(n.pageY-s.top-i.offsetHeight/2-document.documentElement.scrollTop||document.body.scrollTop)+"px",i.style.left=(n.pageX-s.left-i.offsetWidth/2-document.documentElement.scrollLeft||document.body.scrollLeft)+"px"}return i.style.backgroundColor=r.color,i.className="waves-ripple z-active",!1}}return e[a]?e[a].removeHandle=n:e[a]={removeHandle:n},n}var o={bind:function(e,t){e.addEventListener("click",r(e,t),!1)},update:function(e,t){e.removeEventListener("click",e[a].removeHandle,!1),e.addEventListener("click",r(e,t),!1)},unbind:function(e){e.removeEventListener("click",e[a].removeHandle,!1),e[a]=null,delete e[a]}},s=function(e){e.directive("waves",o)};window.Vue&&(window.waves=o,Vue.use(s)),o.install=s;t["a"]=o},"8d41":function(e,t,n){},d31d:function(e,t,n){},eb2f:function(e,t,n){"use strict";n("d31d")},ff9a:function(e,t,n){"use strict";n.r(t);var a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"app-container"},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],key:e.tableKey,staticStyle:{width:"100%"},attrs:{data:e.list,border:"",fit:"","highlight-current-row":"","span-method":e.reduceCell,"row-class-name":e.rowIndexStyle}},[n("el-table-column",{attrs:{label:"nodeName"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[n("span",[e._v(" "+e._s(a.nodeName)+" ")])]}}])}),n("el-table-column",{attrs:{label:"podIp"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[n("span",[e._v(" "+e._s(a.podIp)+" ")])]}}])}),n("el-table-column",{attrs:{label:"namespace"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[n("span",[e._v(" "+e._s(a.metadata_namespace)+" ")])]}}])}),n("el-table-column",{attrs:{label:"容器信息"}},[n("el-table-column",{attrs:{label:"名称"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[n("span",{staticStyle:{"font-weight":"bold"}},[e._v(" "+e._s(a.name)+" ")])]}}])}),n("el-table-column",{attrs:{label:"镜像"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[n("span",[e._v(" "+e._s(a.container_image)+" ")])]}}])}),n("el-table-column",{attrs:{label:"重启次数"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[n("span",[e._v(" "+e._s(a.container_restartCount)+" ")])]}}])}),n("el-table-column",{attrs:{label:"状态","class-name":"status-col"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[n("span",[e._v(" "+e._s(e._f("statusContainerStatusesFilter")(a))+" ")])]}}])}),n("el-table-column",{attrs:{label:"操作",align:"center","class-name":"small-padding fixed-width"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[a.container_ready?n("el-button",{attrs:{size:"mini",type:"success"},on:{click:function(t){return e.goTerminal(a)}}},[e._v(" 进入命令行 ")]):e._e()]}}])})],1),n("el-table-column",{attrs:{label:"创建时间",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[n("span",[e._v(e._s(a.startTime))])]}}])})],1)],1)},r=[],o=(n("b0c0"),n("d81d"),n("d3b7"),n("159b"),n("b64b"),n("4ec3")),s=n("6724"),i=n("ed08"),u={name:"PodListTable",directives:{waves:s["a"]},filters:{statusContainerStatusesFilter:function(e){return e.container_ready&&e.container_started?"running":"other"}},data:function(){return{tableKey:0,list:null,listLoading:!0,listQuery:{id:this.$route.query.id}}},created:function(){void 0!==this.$route.query.id&&0!==parseInt(this.$route.query.id)?this.getList():this.$router.push("/dashboard?tab=clusters")},methods:{getList:function(){var e=this;this.listLoading=!0,Object(o["e"])(this.listQuery).then((function(t){e.list=e.reconstructionData(t.data.items),setTimeout((function(){e.listLoading=!1}),1500)}))},goTerminal:function(e){var t=this;this.listLoading=!0,Object(o["f"])({id:parseInt(this.$route.query.id),pod:e.metadata_name,container:e.name,namespace:e.metadata_namespace}).then((function(e){window.open("/terminal/?token="+e.data.token),t.listLoading=!1})).catch((function(){t.listLoading=!1}))},formatJson:function(e){return this.list.map((function(t){return e.map((function(e){return"timestamp"===e?Object(i["b"])(t[e]):t[e]}))}))},reconstructionData:function(e){var t=[];return e.forEach((function(e,n){var a={};Object.prototype.hasOwnProperty.call(e,"metadata")?Object.keys(e.metadata).forEach((function(t){var n="metadata_"+t;a[n]=e.metadata[t]})):a={metadata_name:"",metadata_namespace:""};var r={rowSpan:1,rowIndex:1,podIp:e.status.podIP,startTime:e.metadata.creationTimestamp,nodeName:e.spec.nodeName};r=Object.assign(r,a),Array.isArray(e.spec.containers)&&e.spec.containers.length>0&&e.spec.containers.forEach((function(n,a){var o=Object.assign({},n,e),s=e.spec.containers.length;0===a?(r.rowSpan=s,r.rowIndex=0,o=Object.assign(o,r)):(r.rowSpan=s,r.rowIndex=1,o=Object.assign(o,r)),Object.prototype.hasOwnProperty.call(e.status,"containerStatuses")&&Array.isArray(e.status.containerStatuses)&&e.status.containerStatuses.length>0&&e.status.containerStatuses.forEach((function(e,t){if(e.name===o.name){var n={};Object.keys(e).forEach((function(t){var a="container_"+t;n[a]=e[t]})),o=Object.assign(o,n)}})),t.push(o)}))})),t},rowIndexStyle:function(e){return 0===e.rowIndex?"":0===e.row.rowIndex?"row_index0_solid":void 0},reduceCell:function(e){var t=e.row,n=(e.column,e.rowIndex,e.columnIndex);return n<4||9===n?1===t.rowIndex?{rowspan:0,colspan:0}:{rowspan:t.rowSpan,colspan:1}:{rowspan:1,colspan:1}}}},c=u,l=(n("eb2f"),n("2877")),d=Object(l["a"])(c,a,r,!1,null,null,null);t["default"]=d.exports}}]);