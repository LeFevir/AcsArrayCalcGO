clear all;

fid = fopen('d:\Downloads Projects\Расчет для статьи\2013-05-22_11-56-50\AbsField_YZ.bin','r');
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
set(gcf, 'color', 'white');
contourf(Y2,X2,Z2,100,'LineStyle','None');
colormap(jet);
%colormap(gray);
colorbar;
set(gca, 'FontSize', 18.0,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 1]);
set(gca,'PlotBoxAspectRatio',[2 1 1]);
%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
%set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');