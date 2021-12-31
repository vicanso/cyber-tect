var k=Object.defineProperty;var q=Object.getOwnPropertySymbols;var R=Object.prototype.hasOwnProperty,C=Object.prototype.propertyIsEnumerable;var L=(e,t,s)=>t in e?k(e,t,{enumerable:!0,configurable:!0,writable:!0,value:s}):e[t]=s,S=(e,t)=>{for(var s in t||(t={}))R.call(t,s)&&L(e,s,t[s]);if(q)for(var s of q(t))C.call(t,s)&&L(e,s,t[s]);return e};import{r as i,R as D,S as E,U,V as A,W as B,X as _,s as b}from"./index.1e068ea3.js";import{h as d,E as j,G as a,D as O,e as l,a as f,r as w,b as H,f as v}from"./ui.c327b114.js";import{u as I,L as N,e as P}from"./naive.650dd42a.js";const X={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},K=a("path",{d:"M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 110c23.196 0 42 18.804 42 42s-18.804 42-42 42s-42-18.804-42-42s18.804-42 42-42zm56 254c0 6.627-5.373 12-12 12h-88c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h12v-64h-12c-6.627 0-12-5.373-12-12v-24c0-6.627 5.373-12 12-12h64c6.627 0 12 5.373 12 12v100h12c6.627 0 12 5.373 12 12v24z",fill:"currentColor"},null,-1);var V=d({name:"InfoCircle",render:function(t,s){return O(),j("svg",X,[K])}});const z="userTracker",F="httpRequest",M="httpError";function T(e,t){return e._time===t._time?0:e._time>t._time?-1:1}const r=l({processing:!1,items:[],count:-1,flux:""}),g=l({processing:!1,items:[]});function $(e){if(e.error){const s=/, message=([\s\S]*)/.exec(e.error);s&&s.length===2&&(e.error=`${s[1]}, ${e.error.replace(s[0],"")}`)}e.result==="0"?e.resultDesc="\u6210\u529F":e.resultDesc="\u5931\u8D25",e.key=e._time,e.createdAt=_(e._time)}const p=l({processing:!1,items:[]}),n=l({processing:!1,items:[],count:-1,flux:""});function G(e){e.key=e._time,e.createdAt=_(e._time)}const u=l({processing:!1,items:[],count:-1,flux:""}),x=l({processing:!1,items:[]}),h=l({processing:!1,items:[]});function Q(e){e.key=e._time,e.createdAt=_(e._time)}async function ne(e){if(!r.processing)try{r.processing=!0;const{data:t}=await i.get(D,{params:e});r.items=t.trackers||[],r.items.sort(T),r.count=t.count||0,r.flux=t.flux||"",r.items.forEach($)}finally{r.processing=!1}}async function ue(){if(!(g.processing||g.items.length!==0))try{g.processing=!0;const e=E.replace(":measurement",z).replace(":tag","action"),{data:t}=await i.get(e);g.items=(t.values||[]).sort()}finally{g.processing=!1}}function ae(){r.items.length=0,r.count=-1,r.flux=""}async function ie(){if(!(p.processing||p.items.length!==0))try{p.processing=!0;const e=E.replace(":measurement",M).replace(":tag","category"),{data:t}=await i.get(e);p.items=(t.values||[]).sort()}finally{p.processing=!1}}async function ce(e){if(!n.processing)try{n.processing=!0;const{data:t}=await i.get(U,{params:e});n.items=t.httpErrors||[],n.count=t.count||0,n.flux=t.flux||"",n.items.forEach(G),n.items.sort(T)}finally{n.processing=!1}}function oe(){n.items.length=0,n.count=-1,n.flux=""}async function le(e){if(!u.processing)try{u.processing=!0;const{data:t}=await i.get(A,{params:e});u.items=t.requests||[],u.count=t.count||0,u.flux=t.flux||"",u.items.forEach(Q),u.items.sort(T)}finally{u.processing=!1}}function fe(){u.items.length=0,u.count=-1,u.flux=""}async function me(){if(!(x.processing||x.items.length!==0))try{x.processing=!0;const e=E.replace(":measurement",F).replace(":tag","service"),{data:t}=await i.get(e);x.items=t.values||[]}finally{x.processing=!1}}async function ge(){if(!(h.processing||h.items.length!==0))try{h.processing=!0;const e=E.replace(":measurement",F).replace(":tag","route"),{data:t}=await i.get(e);h.items=t.values||[]}finally{h.processing=!1}}async function W(e){const t=B.replace(":measurement",e.measurement),{data:s}=await i.get(t,{params:Object.assign({time:e.time},e.tags)});return s}const J={userTrackers:f(r),userTrackerActions:f(g),httpErrors:f(n),httpErrorCategories:f(p),requests:f(u),requestServices:f(x),requestRoutes:f(h)};function pe(){return J}const Y="i18la4si";var Z=d({name:"FluxDetailList",props:{measurement:{type:String,required:!0},time:{type:String,required:!0},tags:{type:Object,required:!0}},setup(e){const t=I(),s=w(!0),c=w([]);return H(async()=>{try{const o=await W({measurement:e.measurement,time:e.time,tags:e.tags}),y=["_measurement","_start","_stop","_time","result","table"];Object.keys(o).forEach(m=>{y.includes(m)||c.value.push({name:m,value:o[m]})})}catch(o){b(t,o)}finally{s.value=!1}}),{processing:s,values:c}},render(){const{processing:e,values:t}=this;if(e)return a("span",null,[v("\u6B63\u5728\u52A0\u8F7D\u4E2D...")]);if(t.length===0)return a("span",null,[v("\u5F88\u62B1\u6B49\uFF0C\u65E0\u7B26\u5408\u8BB0\u5F55")]);const s=t.map(c=>a("li",null,[a("span",{class:"mright5"},[c.name,v(":")]),v(" "),String(c.value)]));return a("ul",{class:Y},[s])}}),xe=d({name:"ExFluxDetail",props:{measurement:{type:String,required:!0},data:{type:Object,required:!0},tagKeys:{type:Array,required:!0}},render(){const{data:e,measurement:t,tagKeys:s}=this.$props,c={trigger:()=>a(P,null,{default:()=>[a(V,null,null)]})},o={};return Object.keys(e).forEach(y=>{if(!s.includes(y))return;const m=e[y];!m||(o[y]=m)}),a(N,{trigger:"hover",placement:"top-end",delay:500,duration:1e3},S({default:()=>[a(Z,{measurement:t,time:e._time,tags:o},null)]},c))}});export{xe as E,ae as a,ne as b,ie as c,oe as d,M as e,ue as f,ce as g,me as h,ge as i,fe as j,F as k,le as l,z as m,pe as u};
