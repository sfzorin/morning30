import math, os
FT="M 8,-6 C 34,-6 50,0 57,9 C 63,13 61,21 51,21 L -8,21 C -18,21 -18,-3 -9,-5 Z"
L={'torso':175,'neck':26,'uarm':116,'farm':90,'thigh':151,'shin':151,'foot':63}
shHW=52; hipHW=36; chHD=26; hipHD=23; HEAD_OFF=64
def nz(v):
    m=math.sqrt(sum(c*c for c in v)) or 1; return tuple(c/m for c in v)
def add(p,d,l): d=nz(d); return tuple(p[i]+d[i]*l for i in range(3))
def addv(p,u,l): return tuple(p[i]+u[i]*l for i in range(3))
def cross(a,b): return (a[1]*b[2]-a[2]*b[1],a[2]*b[0]-a[0]*b[2],a[0]*b[1]-a[1]*b[0])
def dot(a,b): return sum(a[i]*b[i] for i in range(3))
def scl(v,s): return tuple(c*s for c in v)
def clamp(T,S,allowed):
    T=nz(T); S=nz(S); d=dot(S,T); Sp=tuple(S[i]-d*T[i] for i in range(3))
    return T if dot(Sp,allowed)<-1e-6 else S
def project(p,v):
    x,y,z=p
    return ( x,-y,z) if v=='side' else ( z,-y,x) if v=='front' else ( x,z,-y)

# ---- skeleton: all joint positions in 3D ----
def joints(pose):
    pel=pose['pelvis']; spine=nz(pose['spine']); front=nz(pose['front']); lat=nz(cross(spine,front))
    chest=add(pel,spine,L['torso'])
    J={'pel':pel,'chest':chest,'neck':add(chest,spine,L['neck']),'head':add(chest,spine,HEAD_OFF),
       'spine':spine,'front':front,'lat':lat}
    for s,sg in (('L',1),('R',-1)):
        J['hip'+s]=addv(pel,lat,sg*hipHW); J['sh'+s]=addv(chest,lat,sg*shHW)
        ua=pose[s+'uarm']; fa=clamp(ua,pose[s+'farm'],front)
        J['el'+s]=add(J['sh'+s],ua,L['uarm']); J['wr'+s]=add(J['el'+s],fa,L['farm'])
        th=pose[s+'thigh']; sd=clamp(th,pose[s+'shin'],scl(front,-1))
        J['kn'+s]=add(J['hip'+s],th,L['thigh']); J['an'+s]=add(J['kn'+s],sd,L['shin'])
        J['to'+s]=add(J['an'+s],pose[s+'foot'],L['foot'])
    return J

class Fig:
    def __init__(s,view): s.view=view; s.draw=[]; s.pts=[]
    def P(s,p):
        a=project(p,s.view); s.pts.append((a[0],a[1])); return (a[0],a[1]),a[2]
    def Pf(s,p):
        a=project(p,s.view); return (a[0],a[1]),a[2]
    def tube(s,pts3,depth):
        sc=[s.P(p)[0] for p in pts3]
        d=" ".join(("M " if i==0 else "L ")+f"{round(x,1)},{round(y,1)}" for i,(x,y) in enumerate(sc))
        s.draw.append((depth,f'<path d="{d}" fill="none" stroke="#1F2937" stroke-width="40"/><path d="{d}" fill="none" stroke="#E5E7EB" stroke-width="28"/>'))
    def foot(s,an,to,depth):
        (ax,ay),_=s.P(an);(tx,ty),_=s.P(to); ln=math.hypot(tx-ax,ty-ay); s.pts.append((ax,ay+22))
        if ln<16:
            s.draw.append((depth+.05,f'<ellipse cx="{round(ax,1)}" cy="{round(ay+10,1)}" rx="17" ry="12" fill="#E5E7EB" stroke="#1F2937" stroke-width="6"/>')); return
        deg=math.degrees(math.atan2(ty-ay,tx-ax)); mir=' scale(1 -1)' if math.cos(math.radians(deg))<0 else ''
        s.draw.append((depth+.05,f'<use href="#FT" transform="translate({round(ax,1)} {round(ay,1)}) rotate({round(deg,1)}){mir}" fill="#E5E7EB" stroke="#1F2937" stroke-width="6"/>'))
    def head(s,J):
        (cx,cy),dz=s.P(J['head']);(sx,sy),_=s.Pf(J['chest']);(hx,hy),_=s.Pf(J['neck'])
        a=round(math.degrees(math.atan2(hx-sx,-(hy-sy))),1)
        s.draw.append((dz+.2,f'<ellipse cx="{round(cx,1)}" cy="{round(cy,1)}" rx="29" ry="41" fill="#E5E7EB" stroke="#1F2937" stroke-width="6" transform="rotate({a} {round(cx,1)} {round(cy,1)})"/>'))
    def torso(s,J):
        # loft cross-sections between pelvis and chest; radius depends on view (lateral width vs sagittal depth)
        useW = s.view!='side'
        prof=[(0.0,hipHW,hipHD),(0.45,hipHW*0.8,hipHD*0.78),(0.82,shHW*0.86,chHD*0.96),(1.0,shHW,chHD)]
        (px,py),dz=s.Pf(J['pel']);(qx,qy),_=s.Pf(J['chest'])
        dx,dy=qx-px,qy-py; dl=math.hypot(dx,dy) or 1; ux,uy=dx/dl,dy/dl; nx,ny=-uy,ux
        Lp=[];Rp=[]
        for t,w,h in prof:
            (sx,sy),_=s.Pf(addv(J['pel'],J['spine'],t*L['torso'])); r=(w if useW else h)
            Lp.append((sx+nx*r,sy+ny*r)); Rp.append((sx-nx*r,sy-ny*r)); s.pts+=[Lp[-1],Rp[-1]]
        # rounded top cap above shoulders, rounded bottom under pelvis
        topc=(qx+ux*8, qy+uy*8); botc=(px-ux*10, py-uy*10)
        pth="M "+f"{round(Lp[0][0],1)},{round(Lp[0][1],1)}"
        for x,y in Lp[1:]: pth+=f" L {round(x,1)},{round(y,1)}"
        pth+=f" Q {round(topc[0],1)},{round(topc[1],1)} {round(Rp[-1][0],1)},{round(Rp[-1][1],1)}"
        for x,y in reversed(Rp[:-1]): pth+=f" L {round(x,1)},{round(y,1)}"
        pth+=f" Q {round(botc[0],1)},{round(botc[1],1)} {round(Lp[0][0],1)},{round(Lp[0][1],1)} Z"
        s.draw.append((dz,f'<path d="{pth}" fill="#E5E7EB" stroke="#1F2937" stroke-width="6" stroke-linejoin="round"/>'))
    def body(s): return "\n      ".join(x for _,x in sorted(s.draw,key=lambda t:t[0]))

def meat(pose,view):
    J=joints(pose); f=Fig(view)
    for s in ('L','R'):
        f.tube([J['hip'+s],J['kn'+s],J['an'+s]], f.Pf(J['hip'+s])[1]); f.foot(J['an'+s],J['to'+s],f.Pf(J['hip'+s])[1])
    f.torso(J)
    f.tube([J['chest'],J['neck']], f.Pf(J['chest'])[1]-0.2)
    for s in ('L','R'):
        f.tube([J['sh'+s],J['el'+s],J['wr'+s]], f.Pf(J['sh'+s])[1])
    f.head(J)
    return f

def skel(pose,view):
    J=joints(pose); f=Fig(view); seg=[]
    bones=[('pel','chest'),('chest','neck'),('neck','head'),('pel','hipL'),('pel','hipR'),
           ('chest','shL'),('chest','shR'),('hipL','knL'),('knL','anL'),('anL','toL'),
           ('hipR','knR'),('knR','anR'),('anR','toR'),('shL','elL'),('elL','wrL'),('shR','elR'),('elR','wrR')]
    for a,b in bones:
        (ax,ay),_=f.P(J[a]);(bx,by),_=f.P(J[b])
        seg.append(f'<line x1="{round(ax,1)}" y1="{round(ay,1)}" x2="{round(bx,1)}" y2="{round(by,1)}" stroke="#1F2937" stroke-width="3"/>')
    dots=""
    for k in ('pel','chest','neck','hipL','hipR','shL','shR','elL','wrL','elR','wrR','knL','anL','knR','anR'):
        (x,y),_=f.Pf(J[k]); dots+=f'<circle cx="{round(x,1)}" cy="{round(y,1)}" r="6" fill="#2563EB"/>'
    f.draw.append((0,"".join(seg)+dots))
    (hx,hy),_=f.Pf(J['head']); f.draw.append((1,f'<circle cx="{round(hx,1)}" cy="{round(hy,1)}" r="32" fill="none" stroke="#2563EB" stroke-width="3"/>'))
    return f

def emit(idd,title,fig):
    xs=[p[0] for p in fig.pts] or [0]; ys=[p[1] for p in fig.pts] or [0]
    minx,maxx,miny,maxy=min(xs),max(xs),min(ys),max(ys); w=maxx-minx or 1; h=maxy-miny or 1
    s=min(880/w,600/h,1.25); tx=512-s*(minx+maxx)/2; ty=700-s*maxy
    floor=""
    if fig.view!="top":
        fy=round(ty+s*maxy+2,1); x1=round(max(60,tx+s*minx-30),1); x2=round(min(964,tx+s*maxx+30),1)
        floor=f'<line x1="{x1}" y1="{fy}" x2="{x2}" y2="{fy}" stroke="#1F2937" stroke-width="4" stroke-linecap="round" opacity="0.45"/>'
    return f'''<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 768" role="img" aria-labelledby="{idd}t"><title id="{idd}t">{title}</title>
  <defs><path id="FT" d="{FT}"/></defs>
  <g>{floor}</g>
  <g transform="translate({round(tx,1)} {round(ty,1)}) scale({round(s,3)})" stroke-linecap="round" stroke-linejoin="round">
      {fig.body()}
  </g></svg>
'''

D=os.path.dirname(__file__)+"/svg3d"; os.makedirs(D,exist_ok=True)
DN=(0,-1,0)
poses={
 'squat':{'pelvis':(0,300,0),'spine':(0.32,0.95,0),'front':(1,0,0),
   'Lthigh':(0.55,-0.84,0),'Lshin':(-0.05,-1,0),'Lfoot':(1,-0.04,0),'Rthigh':(0.55,-0.84,0),'Rshin':(-0.05,-1,0),'Rfoot':(1,-0.04,0),
   'Luarm':(0.4,-0.85,0),'Lfarm':(0.95,-0.05,0),'Ruarm':(0.4,-0.85,0),'Rfarm':(0.95,-0.05,0)},
 'armsout':{'pelvis':(0,326,0),'spine':(0,1,0),'front':(1,0,0),
   'Lthigh':DN,'Lshin':DN,'Lfoot':(1,-0.04,0),'Rthigh':DN,'Rshin':DN,'Rfoot':(1,-0.04,0),
   'Luarm':(0,0.05,1),'Lfarm':(0,0.05,1),'Ruarm':(0,0.05,-1),'Rfarm':(0,0.05,-1)},
}
open(f"{D}/squat_skel.svg","w").write(emit("a","squat — скелет",skel(poses['squat'],'side')))
open(f"{D}/squat_meat.svg","w").write(emit("b","squat — мясо",meat(poses['squat'],'side')))
open(f"{D}/armsout_skel.svg","w").write(emit("c","armsout — скелет",skel(poses['armsout'],'front')))
open(f"{D}/armsout_meat.svg","w").write(emit("d","armsout — мясо",meat(poses['armsout'],'front')))
print("ok")
