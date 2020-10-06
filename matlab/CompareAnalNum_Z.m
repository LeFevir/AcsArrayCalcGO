clear all;

fid = fopen('d:\Downloads\Calc NEW\Calc\one_compare\2015-03-25_19-42-11\AbsField_Z.bin','r');
isize=fread(fid, 1, 'int64', 'l');

for i=1:isize
        X1(i)=fread(fid, 1, 'float64', 'l')*1000;
        Z1(i)=fread(fid, 1, 'float64', 'l');
end;
fclose(fid);

fid = fopen('d:\Downloads\Calc NEW\Calc\one_compare\2015-03-25_19-19-40\AbsField_Z.bin','r');
isize=fread(fid, 1, 'int64', 'l');

for i=1:isize
        X2(i)=fread(fid, 1, 'float64', 'l')*1000;
        Z2(i)=fread(fid, 1, 'float64', 'l');
end;
fclose(fid);


figure(11); 
set(gcf, 'color', 'white');
plot(X1,Z1,X2,Z2);
%axis([0 180 0 60]);

% figure(2); 
% set(gcf, 'color', 'white');
% plot(X1,Z1,X11,Z11);
% legend('Analytic', 'Numeric');
% %set(h2,'LineColor','k');
% %set(gca,'CLIm',[0 1]);
% %set(gca,'PlotBoxAspectRatio',[1 1 1]);
% %xlabel('y, mm');
% %ylabel('z, mm');
% %axis([-10 10 120 140]);
% %set(gca,'ytick',[ ]);
% %title('Nonlinear 10 W/cm^{2} after 0.76s');
% %title('(b)');
