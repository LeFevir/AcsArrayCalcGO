% fid = fopen('f:\Educ\��������\.�����\Java\Pressure.txt','r');
% A=fscanf(fid,'%f %f', [2 inf]);
% fclose(fid);
clear all;

% ��������� ������ ����������� ���������
% fid = fopen('offElements.txt','r');
% A=fscanf(fid,'%d', inf);
% fclose(fid);


figure (10);
set(gcf, 'color', 'white');




fid = fopen('d:\Downloads Projects\������ ��� ������\������� ��������� �������\array_elements.txt','r');
A=fscanf(fid,'%f %f', [2 inf]);
fclose(fid);

aperture = 0.170*1000;
hole = 0.020*1000;
elementRadius = 0.0035*1000;
% ������ �������� ����������
plotCircle(0, 0, aperture/2, 3, 'w', '');
hold on;
% ��������� ��� �����
plotCircle(0, 0, hole, 3, 'w', '');

numOfElements = length(A);

for i = 1:numOfElements
    posx = A(1,i)*1000;
    posy = A(2,i)*1000;
    
    %plotCircle(posx, posy, elementRadius, 'w', num2str(i-1));
    plotCircle(posx, posy, elementRadius, 2, [0.7 0.7 0.7], '');
    
    %             if find(A==id)
    %                 % ������ ����������� �������
    %                 plotCircle(posx*1000, posy*1000, elementRadius, 'w', num2str(id));
    %             else
    %                 % ������ ���������� �������
    %                 plotCircle(posx*1000, posy*1000, elementRadius, 'r', num2str(id));
    %             end
end
hold off;
set(gca,'PlotBoxAspectRatio',[1 1 1]);
