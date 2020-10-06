clear all;

fid = fopen('d:\Educ\¿ ”—“» ¿\.œ–Œ√¿\GO\AbsField_XZ.txt','r');
isize=fscanf(fid,'%d ',1);
jsize=fscanf(fid,'%d\n',1);


for i=1:isize
    for j=1:jsize
        X3(i,j)=fscanf(fid,'%f ',1)*1000;
        Y3(i,j)=fscanf(fid,'%f ',1)*1000;
        Z3(i,j)=fscanf(fid,'%f\n',1);
    end;
end;
fclose(fid);

figure(1); 
set(gcf, 'color', 'white');
contourf(Y3,X3,Z3,100,'LineStyle','None');
colormap(jet);
colorbar;
set(gca, 'FontSize', 24.0,'FontName','Times New Roman');
%set(h2,'LineColor','k');
%set(gca,'CLIm',[0 1]);
set(gca,'PlotBoxAspectRatio',[2 1 1]);
%xlabel('y, mm');
%ylabel('z, mm');
%axis([-10 10 120 140]);
%set(gca,'ytick',[ ]);
%title('Nonlinear 10 W/cm^{2} after 0.76s');
%title('(b)');