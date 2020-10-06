% fid = fopen('f:\Educ\¿ ”—“» ¿\.œ–Œ√¿\Java\Pressure.txt','r');
% A=fscanf(fid,'%f %f', [2 inf]);
% fclose(fid);

clear all;
fsize = 18.0;

fid = fopen('d:\Educ\¿ ”—“» ¿\.œ–Œ√¿\GO SideInfo 5\2012-09-10_10-56-03\1\AbsField_YZ.bin','r');
isize=fread(fid, 1, 'int64', 'l');
jsize=fread(fid, 1, 'int64', 'l');

for i=1:isize
    for j=1:jsize  
        X2(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Y2(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Z2(i,j)=fread(fid, 1, 'float64', 'l');
    end;
end;
fclose(fid);

figure(2); 
subplot(1,4,[1 2]);
set(gcf, 'color', 'white');
contourf(Y2,X2,Z2,100,'LineStyle','None');
colormap(jet);
%colormap(gray);
%colorbar;
%set(gca, 'FontSize', 24.0,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 20]);
%xlabel('y, mm');
%ylabel('z, mm');
hold on;
plot([131.0, 131.0], [-25.0,25.0], 'w--', 'LineWidth',1);
plot([84.0, 84.0], [-25.0,25.0], 'w--', 'LineWidth',1);
hold off;

%axis([-10 10 120 140]);
%set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');
set(gca,'PlotBoxAspectRatio',[10 4.3 1]);
%colorbar;
set(gca, 'FontSize', fsize,'FontName','Times New Roman');

  
fid = fopen('d:\Educ\¿ ”—“» ¿\.œ–Œ√¿\GO SideInfo 5\2012-09-10_11-05-25\AbsField_XY.bin','r');
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

subplot(1,4,3);
set(gcf, 'color', 'white');
contourf(X1,Y1,Z1,100,'LineStyle','None');
colormap(jet);
%colormap(gray);

set(gca, 'FontSize', fsize,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 20]);
set(gca,'PlotBoxAspectRatio',[1 1 1]);
hold on;
plot([0.0, 0.0], [-25.0,25.0], 'w--', 'LineWidth',1);
plot( [-25.0,25.0], [0.0, 0.0], 'w--', 'LineWidth',1);
hold off;

%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');
 


fid = fopen('d:\Educ\¿ ”—“» ¿\.œ–Œ√¿\GO SideInfo 5\2012-09-10_10-52-10\AbsField_XY.bin','r');
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

subplot(1,4,4);
set(gcf, 'color', 'white');
contourf(X1,Y1,Z1,100,'LineStyle','None');
colormap(jet);
%colormap(gray);
colorbar;
set(gca, 'FontSize', fsize,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 20]);
set(gca,'PlotBoxAspectRatio',[1 1 1]);
hold on;
plot([0.0, 0.0], [-25.0,25.0], 'w--', 'LineWidth',1);
plot( [-25.0,25.0], [0.0, 0.0], 'w--', 'LineWidth',1);
hold off;
%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');
 





