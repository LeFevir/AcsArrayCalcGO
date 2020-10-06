% fid = fopen('f:\Educ\АКУСТИКА\.ПРОГА\Java\Pressure.txt','r');
% A=fscanf(fid,'%f %f', [2 inf]);
% fclose(fid);

clear all;

MAIN_MAX_SPOT_DY = 5.0;
MAIN_MAX_SPOT_DZ = 15.0;

fsize = 16.0;

for t=1:111
fid = fopen(['d:\Educ\АКУСТИКА\.ПРОГА\GO SideInfo 6\2012-09-12_10-33-15 Z\',num2str(t),'\AbsField_YZ.bin'],'r');
isize=fread(fid, 1, 'int64', 'l');
jsize=fread(fid, 1, 'int64', 'l');

X=zeros(isize,jsize);
Y=zeros(isize,jsize);
Z=zeros(isize,jsize);

for i=1:isize
    for j=1:jsize  
        X(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        Y(i,j)=fread(fid, 1, 'float64', 'l')*1000;
        d=fread(fid, 1, 'float64', 'l');
        Z(i,j)=d;
    end;
end;

%FIELD INFO
i_f=fread(fid, 1, 'int64', 'l') + 1; % +1 because in Matlab indexes start from 1
j_f=fread(fid, 1, 'int64', 'l') + 1;
k_f=fread(fid, 1, 'int64', 'l') + 1;

i_m=fread(fid, 1, 'int64', 'l') + 1;
j_m=fread(fid, 1, 'int64', 'l') + 1;
k_m=fread(fid, 1, 'int64', 'l') + 1;

i_s=fread(fid, 1, 'int64', 'l') + 1;
j_s=fread(fid, 1, 'int64', 'l') + 1;
k_s=fread(fid, 1, 'int64', 'l') + 1;

fclose(fid);



h = figure(1);
set(h, 'color', 'white');
contourf(Y,X,Z,30,'LineStyle','None');
%set(gca,'CLIm',[0 1]);
colormap(jet);
set(gca,'PlotBoxAspectRatio',[2 1 1]);
colorbar;
xlabel('z, мм','FontSize', fsize,'FontName','Times New Roman');
ylabel('y, мм','FontSize', fsize,'FontName','Times New Roman');
zlabel('p/p_0','FontSize', fsize,'FontName','Times New Roman');
% 
% xlabel('z, mm','FontSize', fsize,'FontName','Times New Roman');
% ylabel('y, mm','FontSize', fsize,'FontName','Times New Roman');
% zlabel('p/p_0','FontSize', fsize,'FontName','Times New Roman');

hold on;

%Рисуем положение фокуса
plot([130.0, 130.0], [-30.0, 30.0],'Color','w', 'LineWidth',1, 'LineStyle','--');
plot([70.0,200.0], [0.0, 0.0],'Color','w', 'LineWidth',1, 'LineStyle','--');

ratio = Z(j_s, k_s)/Z(j_m, k_m);

if ratio >0.32
    plot(Y(j_s, k_s), X(j_s, k_s),'ko','LineWidth',2,'MarkerSize',17, 'MarkerFaceColor', 'r');
else
    plot(Y(j_s, k_s), X(j_s, k_s),'ko','LineWidth',2,'MarkerSize',7, 'MarkerFaceColor', 'g');
end

rectangle('Position',[Y(j_f, k_f) - MAIN_MAX_SPOT_DZ, X(j_f, k_f) - MAIN_MAX_SPOT_DY, 2*MAIN_MAX_SPOT_DZ, 2*MAIN_MAX_SPOT_DY], 'EdgeColor','w', 'LineWidth',1, 'LineStyle','-');

hold off;

title(['Фокус в ', num2str(X(j_f, k_f)), ', ', num2str(Y(j_f, k_f)),' мм. Поб. макс. в ', num2str(X(j_s, k_s)), ', ', num2str(Y(j_s, k_s)),' мм. Отношение ', num2str(ratio)], 'FontSize', fsize);

set(gca, 'FontSize', fsize,'FontName','Times New Roman');

set(h,'Visible','off');

saveas(h, ['d:\MatlabMovie3\New\pics Z\', num2str(t)], 'tiff')

end;