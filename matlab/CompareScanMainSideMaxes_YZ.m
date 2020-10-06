clear all;
fsize=24;

fid = fopen('d:\Educ\¿ ”—“» ¿\.Current\–‡Ò˜ÂÚ ‡Í ÔÓÎˇ Â¯ÂÚÍË\Scanning Maxes\Hand\AbsField_YZ.bin','r');
isize=fread(fid, 1, 'int64', 'l');
jsize=fread(fid, 1, 'int64', 'l');

for i=1:isize
    for j=1:jsize  
        X2(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Y2(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Z2(i,j) = fread(fid, 1, 'float64', 'l');
    end;
end;
fclose(fid);

figure(2); 
set(gcf, 'color', 'white');
[C h] = contour(Y2, X2, Z2,'b','LevelList',[0.5 0.7 0.9],'LineWidth',2);
hold on;
clabel(C,h,'FontSize',fsize-10,'FontName','Times New Roman','LabelSpacing',152);
set(gca, 'FontSize', fsize,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 1]);
set(gca,'PlotBoxAspectRatio',[2 1 1]);
%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
%set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');


fid = fopen('d:\Educ\¿ ”—“» ¿\.œ–Œ√¿\GO SideScan YZ\2012-09-11_15-48-38\AbsField_YZ.bin','r');
isize=fread(fid, 1, 'int64', 'l');
jsize=fread(fid, 1, 'int64', 'l');

for i=1:isize
    for j=1:jsize  
        X3(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Y3(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Z3(i,j) = fread(fid, 1, 'float64', 'l');
    end;
end;
fclose(fid);

figure(3); 
set(gcf, 'color', 'white');
[C h] = contour(Y3, X3, Z3,'r','LevelList',[0.32],'LineWidth',4);
hold off;
clabel(C,h,'FontSize',fsize-10,'FontName','Times New Roman','LabelSpacing',152);
%set(gca, 'FontSize', fsize,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 1]);
%set(gca,'PlotBoxAspectRatio',[2 1 1]);
%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
%set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');