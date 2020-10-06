% fid = fopen('f:\Educ\АКУСТИКА\.ПРОГА\Java\Pressure.txt','r');
% A=fscanf(fid,'%f %f', [2 inf]);
% fclose(fid);

clear all;
fsize = 18.0;


% fid = fopen('d:\Educ\АКУСТИКА\.Current\Расчет ак поля решетки\Scanning Maxes\Hand\AbsField_YZ.bin','r');
% isize=fread(fid, 1, 'int64', 'l');
% jsize=fread(fid, 1, 'int64', 'l');
% 
% for i=1:isize
%     for j=1:jsize  
%         X22(i,j)=fread(fid, 1, 'float64', 'l')*1000;
%         Y22(i,j)=fread(fid, 1, 'float64', 'l')*1000;
%         Z22(i,j)=fread(fid, 1, 'float64', 'l');
%     end;
% end;
% fclose(fid);
% 
% figure(2); 
% subplot(1,3,[1 2]);
% set(gcf, 'color', 'white');
% [C h] = contour(Y22, X22, Z22, 'k','LevelList',[0.5 0.7 0.9],'LineWidth',2);
% %colormap jet;
% hold on;
% % clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);
% set(gca,'PlotBoxAspectRatio',[10 4.3 1]);
% set(gca, 'FontSize', fsize,'FontName','Times New Roman');
% 
% 
% 

fid = fopen('d:\Downloads\Calc NEW\Calc\main_max\2015-03-26_11-07-56\AbsField_YZ.bin','r');

isize=fread(fid, 1, 'int64', 'l');
jsize=fread(fid, 1, 'int64', 'l');

for i=1:isize
    for j=1:jsize  
        X11(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Y11(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        temp = fread(fid, 1, 'float64', 'l')
        Z11(i,j)=temp*temp;
    end;
end;
fclose(fid);

figure(2);
%subplot(1,3,3);
set(gcf, 'color', 'white');
[C h] = contour(Y11, X11, Z11,'LevelList',[0.5],'LineWidth',2);
%hold on;
clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);
set(gca, 'FontSize', fsize,'FontName','Times New Roman');
set(gca,'PlotBoxAspectRatio',[200 180 1]);
%set(gca,'ytick',[ ]);


hold on;


fid = fopen('d:\Downloads\Calc NEW\Calc\side_max\2015-03-26_10-34-14\AbsField_YZ.bin','r');
isize=fread(fid, 1, 'int64', 'l');
jsize=fread(fid, 1, 'int64', 'l');

for i=1:isize
    for j=1:jsize  
        X3(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Y3(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        d=fread(fid, 1, 'float64', 'l');
        Z3(i,j)=d*d;
    end;
end;
fclose(fid);

figure(2);
%subplot(1,3,[1 2]);
set(gcf, 'color', 'white');
[C h] = contour(Y3, X3, Z3,'LevelList',[0.1],'LineWidth',3);
clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',252);
%set(gca,'PlotBoxAspectRatio',[10 4.3 1]);
set(gca, 'FontSize', fsize,'FontName','Times New Roman');
xlabel('z, mm');
ylabel('y, mm');


hold off;


% fid = fopen('d:\Educ\АКУСТИКА\.ПРОГА\GO SideScan YZ\2012-09-11_15-48-38\AbsField_YZ.bin','r');
% isize=fread(fid, 1, 'int64', 'l');
% jsize=fread(fid, 1, 'int64', 'l');
% 
% for i=1:isize
%     for j=1:jsize  
%         X2(i,j)=fread(fid, 1, 'float64', 'l')*1000;
%         Y2(i,j)=fread(fid, 1, 'float64', 'l')*1000;
%         d=fread(fid, 1, 'float64', 'l');
%         Z2(i,j)=d;
%     end;
% end;
% fclose(fid);
% 
% figure(2); 
% subplot(1,3,[1 2]);
% set(gcf, 'color', 'white');
% [C h] = contour(Y2, X2, Z2,'r','LevelList',[0.32],'LineWidth',3);
% % clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);
% hold off;
% set(gca,'PlotBoxAspectRatio',[10 4.3 1]);
% set(gca, 'FontSize', fsize,'FontName','Times New Roman');
% 
% 
% fid = fopen('d:\Educ\АКУСТИКА\.Current\Расчет ак поля решетки\Scanning Maxes\Hand\AbsField_XY.bin','r');
% 
% isize=fread(fid, 1, 'int64', 'l');
% jsize=fread(fid, 1, 'int64', 'l');
% 
% for i=1:isize
%     for j=1:jsize  
%         X11(i,j)=fread(fid, 1, 'float64', 'l')*1000;
%         Y11(i,j)=fread(fid, 1, 'float64', 'l')*1000;
%         Z11(i,j)=fread(fid, 1, 'float64', 'l');
%     end;
% end;
% fclose(fid);
% 
% subplot(1,3,3);
% set(gcf, 'color', 'white');
% [C h] = contour(Y11, X11, Z11,'k','LevelList',[0.5 0.7 0.9],'LineWidth',2);
% hold on;
% % clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);
% set(gca, 'FontSize', fsize,'FontName','Times New Roman');
% set(gca,'PlotBoxAspectRatio',[1 1 1]);
% set(gca,'ytick',[ ]);
% 
% 
% fid = fopen('d:\Educ\АКУСТИКА\.Current\Расчет ак поля решетки\Scanning Side Maxes\Hand\PressureXY.bin','r');
% isize=fread(fid, 1, 'uint', 'b');
% jsize=fread(fid, 1, 'uint', 'b');
% 
% for i=1:isize
%     for j=1:jsize  
%         X1(i,j)=fread(fid, 1, 'double', 'b')*1000;
%         Y1(i,j)=fread(fid, 1, 'double', 'b')*1000;
%         d=fread(fid, 1, 'double', 'b');
%         Z1(i,j)=d;
%     end;
% end;
% fclose(fid);
% 
% 
% subplot(1,3,3);
% set(gcf, 'color', 'white');
% [C h] = contour(Y1, X1, Z1,'g','LevelList',[0.32],'LineWidth',3);
% % clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);
% set(gca, 'FontSize', fsize,'FontName','Times New Roman');
% set(gca,'PlotBoxAspectRatio',[1 1 1]);
% set(gca,'ytick',[ ]);
% 
% 
% 
% fid = fopen('d:\Educ\АКУСТИКА\.ПРОГА\GO SideScan XY\2012-12-18_17-26-24\AbsField_XY.bin','r');
% isize=fread(fid, 1, 'int64', 'l');
% jsize=fread(fid, 1, 'int64', 'l');
% 
% for i=1:isize
%     for j=1:jsize  
%         X1(i,j)=fread(fid, 1, 'float64', 'l')*1000;
%         Y1(i,j)=fread(fid, 1, 'float64', 'l')*1000;
%         d=fread(fid, 1, 'float64', 'l');
%         Z1(i,j)=d;
%     end;
% end;
% fclose(fid);
% 
% subplot(1,3,3);
% set(gcf, 'color', 'white');
% [C h] = contour(Y1, X1, Z1,'r','LevelList',[0.32],'LineWidth',3);
% hold off;
% % clabel(C,h,'FontSize',fsize-5,'FontName','Times New Roman','LabelSpacing',152);
% set(gca, 'FontSize', fsize,'FontName','Times New Roman');
% set(gca,'PlotBoxAspectRatio',[1 1 1]);
% set(gca,'ytick',[ ]);
% 
% %legend('уменьшение основного максимума','доля побочного максимума')




