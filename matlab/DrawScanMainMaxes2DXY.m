clear all;
fsize=24

fid = fopen('d:\Educ\¿ ”—“» ¿\.œ–Œ√¿\GO\AbsField_XY.bin','r');
isize=fread(fid, 1, 'int64', 'l');
jsize=fread(fid, 1, 'int64', 'l');

for i=1:isize
    for j=1:jsize  
        X1(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Y1(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Z1(i,j)=fread(fid, 1, 'float64', 'l');
    end;
end;
fclose(fid);

figure(1); 
set(gcf, 'color', 'white');
[C h] = contour(Y1, X1, Z1,'LevelList',[0.5 0.7 0.9],'LineWidth',2);
clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);
set(gca, 'FontSize', fsize,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 1]);
set(gca,'PlotBoxAspectRatio',[1 1 1]);
%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
%set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');