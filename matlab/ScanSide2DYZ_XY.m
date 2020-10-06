% fid = fopen('f:\Educ\АКУСТИКА\.ПРОГА\Java\Pressure.txt','r');
% A=fscanf(fid,'%f %f', [2 inf]);
% fclose(fid);

clear all;
fsize = 24.0;

fid = fopen('d:\Educ\АКУСТИКА\.Current\Расчет ак поля решетки\Scanning Side Maxes\Hand\PressureYZ.bin','r');
isize=fread(fid, 1, 'uint', 'b');
jsize=fread(fid, 1, 'uint', 'b');

for i=1:isize
    for j=1:jsize  
        X2(i,j)=fread(fid, 1, 'double', 'b')*1000;
        Y2(i,j)=fread(fid, 1, 'double', 'b')*1000;
        d=fread(fid, 1, 'double', 'b');
        Z2(i,j)=d;
    end;
end;
fclose(fid);


figure(2); 
subplot(1,3,[1 2]);
set(gcf, 'color', 'white');
[C h] = contour(Y2, X2, Z2,'LevelList',[0.32],'LineWidth',3);
%clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);
%colormap(gray);
%colorbar;
%set(gca, 'FontSize', 24.0,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 1]);
%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
%set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');
set(gca,'PlotBoxAspectRatio',[10 4.3 1]);
%colorbar;
set(gca, 'FontSize', fsize,'FontName','Times New Roman');

  
fid = fopen('d:\Educ\АКУСТИКА\.Current\Расчет ак поля решетки\Scanning Side Maxes\Hand\PressureXY.bin','r');
isize=fread(fid, 1, 'uint', 'b');
jsize=fread(fid, 1, 'uint', 'b');

for i=1:isize
    for j=1:jsize  
        X1(i,j)=fread(fid, 1, 'double', 'b')*1000;
        Y1(i,j)=fread(fid, 1, 'double', 'b')*1000;
        d=fread(fid, 1, 'double', 'b');
        Z1(i,j)=d;
    end;
end;
fclose(fid);

subplot(1,3,3);
set(gcf, 'color', 'white');
[C h] = contour(Y1, X1, Z1,'LevelList',[0.32],'LineWidth',3);
%clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);

%colormap(gray);

set(gca, 'FontSize', fsize,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 1]);
set(gca,'PlotBoxAspectRatio',[1 1 1]);
%colorbar;
%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');
 





